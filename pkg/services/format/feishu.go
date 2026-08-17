package format

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"gomessage/pkg/models"
)

const FeishuMaxPayloadBytes = 20 * 1024

type feishuMarkdown struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

type feishuCard struct {
	Schema string `json:"schema"`
	Header struct {
		Title struct {
			Tag     string `json:"tag"`
			Content string `json:"content"`
		} `json:"title"`
		Template string `json:"template"`
	} `json:"header"`
	Body struct {
		Elements []feishuMarkdown `json:"elements"`
	} `json:"body"`
}

type feishuMessage struct {
	MsgType string     `json:"msg_type"`
	Card    feishuCard `json:"card"`
}

// PackFeishuMessage renders a Feishu Card JSON 2.0 interactive message.
func PackFeishuMessage(userConfigInfo *models.Client, message string) interface{} {
	var keyword, titleColor string
	if userConfigInfo != nil && userConfigInfo.ExtendFeishu != nil {
		keyword = userConfigInfo.ExtendFeishu.RobotKeyword
		titleColor = userConfigInfo.ExtendFeishu.TitleColor
	}

	payload := feishuMessage{MsgType: "interactive"}
	payload.Card.Schema = "2.0"
	payload.Card.Header.Title.Tag = "plain_text"
	payload.Card.Header.Title.Content = keyword
	payload.Card.Header.Template = titleColor
	payload.Card.Body.Elements = []feishuMarkdown{{
		Tag:     "markdown",
		Content: message,
	}}
	return payload
}

// GenerateFeishuSignature generates the signature required by a signed custom bot.
func GenerateFeishuSignature(timestamp int64, secret string) string {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	mac := hmac.New(sha256.New, []byte(stringToSign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// MarshalFeishuMessage adds a fresh signature when configured and validates
// Feishu's 20 KB custom-bot request limit.
func MarshalFeishuMessage(message any, secret string, timestamp int64) ([]byte, error) {
	payload, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("marshal Feishu message: %w", err)
	}

	var object map[string]json.RawMessage
	if err := json.Unmarshal(payload, &object); err != nil {
		return nil, fmt.Errorf("Feishu message must be a JSON object: %w", err)
	}
	if object == nil {
		return nil, errors.New("Feishu message must be a JSON object")
	}

	if secret != "" {
		timestampJSON, _ := json.Marshal(strconv.FormatInt(timestamp, 10))
		signJSON, _ := json.Marshal(GenerateFeishuSignature(timestamp, secret))
		object["timestamp"] = timestampJSON
		object["sign"] = signJSON
	}

	payload, err = json.Marshal(object)
	if err != nil {
		return nil, fmt.Errorf("marshal Feishu message: %w", err)
	}
	if len(payload) > FeishuMaxPayloadBytes {
		return nil, fmt.Errorf("Feishu message is %d bytes, exceeds %d-byte limit", len(payload), FeishuMaxPayloadBytes)
	}
	return payload, nil
}

// PackFeishuMessages splits oversized Markdown into multiple valid cards. Rune
// boundaries are preserved and concatenating each card's content restores the
// original message exactly.
func PackFeishuMessages(client *models.Client, message string) ([]any, error) {
	const sizeCheckSecret = "size-check"
	const sizeCheckTimestamp = int64(9999999999)

	if _, err := MarshalFeishuMessage(PackFeishuMessage(client, message), sizeCheckSecret, sizeCheckTimestamp); err == nil {
		return []any{PackFeishuMessage(client, message)}, nil
	}

	runes := []rune(message)
	if len(runes) == 0 {
		return nil, errors.New("Feishu card metadata exceeds the request size limit")
	}

	messages := make([]any, 0, 2)
	for start := 0; start < len(runes); {
		low, high := start+1, len(runes)
		best := start
		for low <= high {
			middle := low + (high-low)/2
			candidate := PackFeishuMessage(client, string(runes[start:middle]))
			if _, err := MarshalFeishuMessage(candidate, sizeCheckSecret, sizeCheckTimestamp); err == nil {
				best = middle
				low = middle + 1
			} else {
				high = middle - 1
			}
		}
		if best == start {
			return nil, errors.New("Feishu card metadata leaves no room for message content")
		}
		messages = append(messages, PackFeishuMessage(client, string(runes[start:best])))
		start = best
	}
	return messages, nil
}
