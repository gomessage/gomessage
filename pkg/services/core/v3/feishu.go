package v3

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"gomessage/pkg/models"
	v1 "gomessage/pkg/services/core/v1"
	"gomessage/pkg/services/format"
	"gomessage/pkg/utils"
	"gomessage/pkg/utils/log/loggers"
)

const (
	feishuRequestTimeout = 10 * time.Second
	feishuMaxAttempts    = 3
	feishuRateLimitCode  = 11232
)

type FeishuAPIError struct {
	HTTPStatus int
	Code       int
	Message    string
}

func (e *FeishuAPIError) Error() string {
	if e.Code != 0 {
		return fmt.Sprintf("Feishu API error: code=%d message=%s", e.Code, e.Message)
	}
	return fmt.Sprintf("Feishu HTTP error: status=%d message=%s", e.HTTPStatus, e.Message)
}

type ClientActionFeishu struct {
	Client     *models.Client
	HTTPClient *http.Client
	Now        func() time.Time
	Sleep      func(time.Duration)
}

func (c *ClientActionFeishu) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var content []string
	if isMerge {
		content = []string{v1.MessageJoint(contentList, utils.VarFeishu)}
	} else {
		content = contentList
	}

	var messages []any
	for _, item := range content {
		packed, err := format.PackFeishuMessages(client, item)
		if err != nil {
			loggers.DefaultLogger.WithError(err).Error("飞书消息渲染失败")
			continue
		}
		messages = append(messages, packed...)
	}
	return messages
}

func (c *ClientActionFeishu) PushMessages(messages []any) (successCount int, failureCount int) {
	webhookURL, secret, err := c.selectRobot()
	if err != nil {
		loggers.DefaultLogger.WithError(err).Error("飞书机器人配置无效")
		return 0, len(messages)
	}

	for _, message := range messages {
		if err := c.pushMessage(message, webhookURL, secret); err != nil {
			failureCount++
			loggers.DefaultLogger.WithFields(logrus.Fields{
				"client_id": c.clientID(),
				"endpoint":  feishuEndpointForLog(webhookURL),
			}).WithError(err).Error("飞书消息推送失败")
		} else {
			successCount++
		}
	}
	return successCount, failureCount
}

func (c *ClientActionFeishu) selectRobot() (string, string, error) {
	if c.Client == nil || c.Client.ExtendFeishu == nil {
		return "", "", errors.New("missing Feishu client configuration")
	}

	urls := c.Client.ExtendFeishu.RobotUrlRandomList
	validIndexes := make([]int, 0, len(urls))
	for index, webhookURL := range urls {
		if strings.TrimSpace(webhookURL) != "" {
			validIndexes = append(validIndexes, index)
		}
	}
	if len(validIndexes) == 0 {
		return "", "", errors.New("missing Feishu webhook URL")
	}

	index := validIndexes[rand.Intn(len(validIndexes))]
	secret := ""
	if index < len(c.Client.ExtendFeishu.RobotSecretRandomList) {
		secret = c.Client.ExtendFeishu.RobotSecretRandomList[index]
	}
	webhookURL := strings.TrimSpace(urls[index])
	parsed, err := url.Parse(webhookURL)
	if err != nil || parsed.Host == "" || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		return "", "", errors.New("invalid Feishu webhook URL")
	}
	return webhookURL, secret, nil
}

func (c *ClientActionFeishu) pushMessage(message any, webhookURL, secret string) error {
	client := c.HTTPClient
	if client == nil {
		client = &http.Client{Timeout: feishuRequestTimeout}
	}

	for attempt := 1; attempt <= feishuMaxAttempts; attempt++ {
		now := feishuLimiterFor(webhookURL).wait(c.now, c.sleep)

		payload, err := format.MarshalFeishuMessage(message, secret, now.Unix())
		if err != nil {
			return err
		}
		request, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewReader(payload))
		if err != nil {
			return errors.New("create Feishu request failed")
		}
		request.Header.Set("Content-Type", "application/json; charset=utf-8")

		response, err := client.Do(request)
		if err != nil {
			if attempt < feishuMaxAttempts {
				c.sleep(retryDelay(attempt, false))
				continue
			}
			return safeFeishuTransportError(err)
		}

		body, readErr := io.ReadAll(io.LimitReader(response.Body, 1<<20))
		closeErr := response.Body.Close()
		if readErr != nil {
			return fmt.Errorf("read Feishu response: %w", readErr)
		}
		if closeErr != nil {
			loggers.DefaultLogger.WithError(closeErr).Warn("关闭飞书响应失败")
		}

		apiErr := parseFeishuResponse(response.StatusCode, body)
		logFeishuAttempt(c.clientID(), webhookURL, response.StatusCode, apiErr, attempt)
		if apiErr == nil {
			return nil
		}
		if attempt < feishuMaxAttempts && isRetryableFeishuError(apiErr) {
			c.sleep(retryDelay(attempt, apiErr.Code == feishuRateLimitCode || apiErr.HTTPStatus == http.StatusTooManyRequests))
			continue
		}
		return apiErr
	}
	return errors.New("Feishu request attempts exhausted")
}

func safeFeishuTransportError(err error) error {
	if errors.Is(err, context.DeadlineExceeded) {
		return errors.New("send Feishu request timed out")
	}
	var networkError net.Error
	if errors.As(err, &networkError) {
		return fmt.Errorf("send Feishu request failed: network error (timeout=%t)", networkError.Timeout())
	}
	return errors.New("send Feishu request failed")
}

type feishuResponse struct {
	Code          *int   `json:"code"`
	Message       string `json:"msg"`
	StatusCode    *int   `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
}

func parseFeishuResponse(httpStatus int, body []byte) *FeishuAPIError {
	var result feishuResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return &FeishuAPIError{HTTPStatus: httpStatus, Message: "invalid JSON response"}
	}

	message := result.Message
	if message == "" {
		message = result.StatusMessage
	}
	if httpStatus < http.StatusOK || httpStatus >= http.StatusMultipleChoices {
		code := 0
		if result.Code != nil {
			code = *result.Code
		} else if result.StatusCode != nil {
			code = *result.StatusCode
		}
		return &FeishuAPIError{HTTPStatus: httpStatus, Code: code, Message: message}
	}
	if result.Code != nil {
		if *result.Code == 0 {
			return nil
		}
		return &FeishuAPIError{HTTPStatus: httpStatus, Code: *result.Code, Message: message}
	}
	if result.StatusCode != nil {
		if *result.StatusCode == 0 {
			return nil
		}
		return &FeishuAPIError{HTTPStatus: httpStatus, Code: *result.StatusCode, Message: message}
	}
	return &FeishuAPIError{HTTPStatus: httpStatus, Message: "missing business status code"}
}

func isRetryableFeishuError(err *FeishuAPIError) bool {
	return err.Code == feishuRateLimitCode ||
		err.HTTPStatus == http.StatusTooManyRequests ||
		err.HTTPStatus >= http.StatusInternalServerError
}

func retryDelay(attempt int, rateLimited bool) time.Duration {
	if rateLimited {
		return time.Duration(attempt) * time.Second
	}
	return time.Duration(1<<(attempt-1)) * 200 * time.Millisecond
}

func (c *ClientActionFeishu) now() time.Time {
	if c.Now != nil {
		return c.Now()
	}
	return time.Now()
}

func (c *ClientActionFeishu) sleep(duration time.Duration) {
	if c.Sleep != nil {
		c.Sleep(duration)
		return
	}
	time.Sleep(duration)
}

func (c *ClientActionFeishu) clientID() int {
	if c.Client == nil {
		return 0
	}
	return c.Client.ID
}

func feishuEndpointForLog(webhookURL string) string {
	parsed, err := url.Parse(webhookURL)
	if err != nil || parsed.Host == "" {
		return "invalid"
	}
	return parsed.Scheme + "://" + parsed.Host
}

func logFeishuAttempt(clientID int, webhookURL string, status int, apiErr *FeishuAPIError, attempt int) {
	fields := logrus.Fields{
		"client_id":       clientID,
		"endpoint":        feishuEndpointForLog(webhookURL),
		"response_status": status,
		"attempt":         attempt,
	}
	if apiErr != nil {
		fields["business_code"] = apiErr.Code
		fields["response_message"] = apiErr.Message
		loggers.PushLogger.WithFields(fields).Warn("飞书推送响应失败")
		return
	}
	loggers.PushLogger.WithFields(fields).Info("飞书推送成功")
}

type feishuRateLimiter struct {
	mu      sync.Mutex
	history []time.Time
}

var feishuRateLimiters sync.Map

func feishuLimiterFor(webhookURL string) *feishuRateLimiter {
	limiter, _ := feishuRateLimiters.LoadOrStore(webhookURL, &feishuRateLimiter{})
	return limiter.(*feishuRateLimiter)
}

func (l *feishuRateLimiter) wait(nowFunc func() time.Time, sleep func(time.Duration)) time.Time {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := nowFunc()
	minuteAgo := now.Add(-time.Minute)
	firstValid := 0
	for firstValid < len(l.history) && !l.history[firstValid].After(minuteAgo) {
		firstValid++
	}
	l.history = l.history[firstValid:]

	var delay time.Duration
	if len(l.history) >= 5 {
		if candidate := l.history[len(l.history)-5].Add(time.Second).Sub(now); candidate > delay {
			delay = candidate
		}
	}
	if len(l.history) >= 100 {
		if candidate := l.history[len(l.history)-100].Add(time.Minute).Sub(now); candidate > delay {
			delay = candidate
		}
	}
	if delay > 0 {
		sleep(delay)
		now = now.Add(delay)
	}
	l.history = append(l.history, now)
	return now
}
