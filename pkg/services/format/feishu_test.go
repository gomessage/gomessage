package format

import (
	"encoding/json"
	"strings"
	"testing"

	"gomessage/pkg/models"
)

func testFeishuClient() *models.Client {
	return &models.Client{
		ExtendFeishu: &models.Feishu{
			RobotKeyword: "生产告警",
			TitleColor:   "red",
		},
	}
}

func TestPackFeishuMessageUsesCardJSON2(t *testing.T) {
	payload, err := json.Marshal(PackFeishuMessage(testFeishuClient(), "**服务异常**"))
	if err != nil {
		t.Fatalf("marshal message: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal message: %v", err)
	}

	card := got["card"].(map[string]any)
	if card["schema"] != "2.0" {
		t.Fatalf("card schema = %v, want 2.0", card["schema"])
	}
	if _, exists := card["elements"]; exists {
		t.Fatal("card JSON 2.0 must not use top-level elements")
	}

	body := card["body"].(map[string]any)
	elements := body["elements"].([]any)
	markdown := elements[0].(map[string]any)
	if markdown["content"] != "**服务异常**" {
		t.Fatalf("markdown content = %v", markdown["content"])
	}
	if _, exists := markdown["Content"]; exists {
		t.Fatal("markdown content key must be lowercase")
	}
}

func TestGenerateFeishuSignature(t *testing.T) {
	const want = "l1N0gAcBjdwBvGm1xMjOF0XSyaLRpR7tuO5dHfhAYc8="
	if got := GenerateFeishuSignature(1599360473, "demo"); got != want {
		t.Fatalf("signature = %q, want %q", got, want)
	}
}

func TestMarshalFeishuMessageAddsSignature(t *testing.T) {
	payload, err := MarshalFeishuMessage(
		PackFeishuMessage(testFeishuClient(), "告警内容"),
		"demo",
		1599360473,
	)
	if err != nil {
		t.Fatalf("marshal signed message: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unmarshal signed message: %v", err)
	}
	if got["timestamp"] != "1599360473" {
		t.Fatalf("timestamp = %v", got["timestamp"])
	}
	if got["sign"] != "l1N0gAcBjdwBvGm1xMjOF0XSyaLRpR7tuO5dHfhAYc8=" {
		t.Fatalf("sign = %v", got["sign"])
	}
}

func TestPackFeishuMessagesSplitsOversizePayload(t *testing.T) {
	content := strings.Repeat("<告警>&\n", 4000)
	messages, err := PackFeishuMessages(testFeishuClient(), content)
	if err != nil {
		t.Fatalf("pack messages: %v", err)
	}
	if len(messages) < 2 {
		t.Fatalf("message count = %d, want split payload", len(messages))
	}

	var rebuilt strings.Builder
	for i, message := range messages {
		payload, err := MarshalFeishuMessage(message, "demo", 1599360473)
		if err != nil {
			t.Fatalf("marshal message %d: %v", i, err)
		}
		if len(payload) > FeishuMaxPayloadBytes {
			t.Fatalf("message %d payload size = %d, limit = %d", i, len(payload), FeishuMaxPayloadBytes)
		}

		var got map[string]any
		if err := json.Unmarshal(payload, &got); err != nil {
			t.Fatalf("unmarshal message %d: %v", i, err)
		}
		card := got["card"].(map[string]any)
		body := card["body"].(map[string]any)
		elements := body["elements"].([]any)
		markdown := elements[0].(map[string]any)
		rebuilt.WriteString(markdown["content"].(string))
	}

	if rebuilt.String() != content {
		t.Fatal("split messages did not preserve the original content")
	}
}
