package v3

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"gomessage/pkg/models"
	"gomessage/pkg/services/format"
)

func feishuTestClient(webhookURL, secret string) *models.Client {
	return &models.Client{
		ExtendFeishu: &models.Feishu{
			RobotKeyword:          "生产告警",
			TitleColor:            "red",
			RobotUrlRandomList:    []string{webhookURL},
			RobotSecretRandomList: []string{secret},
		},
	}
}

func TestClientActionFeishuPushMessagesSuccess(t *testing.T) {
	var received map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&received); err != nil {
			t.Errorf("decode request: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":0,"msg":"success"}`))
	}))
	defer server.Close()

	client := feishuTestClient(server.URL, "demo")
	action := &ClientActionFeishu{
		Client:     client,
		HTTPClient: server.Client(),
		Now:        func() time.Time { return time.Unix(1599360473, 0) },
	}
	messages := []any{format.PackFeishuMessage(client, "告警内容")}

	success, failure := action.PushMessages(messages)
	if success != 1 || failure != 0 {
		t.Fatalf("success=%d failure=%d, want 1/0", success, failure)
	}
	if received["timestamp"] != "1599360473" {
		t.Fatalf("timestamp = %v", received["timestamp"])
	}
	if received["sign"] != "l1N0gAcBjdwBvGm1xMjOF0XSyaLRpR7tuO5dHfhAYc8=" {
		t.Fatalf("sign = %v", received["sign"])
	}
}

func TestClientActionFeishuRejectsBusinessError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":19021,"msg":"sign match fail"}`))
	}))
	defer server.Close()

	client := feishuTestClient(server.URL, "demo")
	action := &ClientActionFeishu{
		Client:     client,
		HTTPClient: server.Client(),
		Now:        func() time.Time { return time.Unix(1599360473, 0) },
	}

	success, failure := action.PushMessages([]any{format.PackFeishuMessage(client, "告警内容")})
	if success != 0 || failure != 1 {
		t.Fatalf("success=%d failure=%d, want 0/1", success, failure)
	}
}

func TestClientActionFeishuRetriesRateLimit(t *testing.T) {
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if requests.Add(1) == 1 {
			_, _ = w.Write([]byte(`{"code":11232,"msg":"rate limited"}`))
			return
		}
		_, _ = w.Write([]byte(`{"code":0,"msg":"success"}`))
	}))
	defer server.Close()

	client := feishuTestClient(server.URL, "")
	action := &ClientActionFeishu{
		Client:     client,
		HTTPClient: server.Client(),
		Sleep:      func(time.Duration) {},
	}

	success, failure := action.PushMessages([]any{format.PackFeishuMessage(client, "告警内容")})
	if success != 1 || failure != 0 {
		t.Fatalf("success=%d failure=%d, want 1/0", success, failure)
	}
	if got := requests.Load(); got != 2 {
		t.Fatalf("request count = %d, want 2", got)
	}
}

func TestClientActionFeishuRejectsMissingWebhook(t *testing.T) {
	client := feishuTestClient("", "")
	client.ExtendFeishu.RobotUrlRandomList = nil
	action := &ClientActionFeishu{Client: client}

	success, failure := action.PushMessages([]any{format.PackFeishuMessage(client, "告警内容")})
	if success != 0 || failure != 1 {
		t.Fatalf("success=%d failure=%d, want 0/1", success, failure)
	}
}

func TestParseFeishuResponseSupportsHistoricalSuccess(t *testing.T) {
	if err := parseFeishuResponse(http.StatusOK, []byte(`{"StatusCode":0,"StatusMessage":"success"}`)); err != nil {
		t.Fatalf("historical success response rejected: %v", err)
	}
}

func TestParseFeishuResponseRejectsMissingBusinessCode(t *testing.T) {
	if err := parseFeishuResponse(http.StatusOK, []byte(`{"msg":"success"}`)); err == nil {
		t.Fatal("response without a business status code must fail")
	}
}

func TestFeishuRateLimiterHonorsSecondLimit(t *testing.T) {
	limiter := &feishuRateLimiter{}
	base := time.Unix(1700000000, 0)
	var slept time.Duration
	for index := 0; index < 6; index++ {
		limiter.wait(func() time.Time { return base }, func(duration time.Duration) {
			slept += duration
		})
	}
	if slept != time.Second {
		t.Fatalf("rate-limit sleep = %s, want 1s", slept)
	}
}

func TestFeishuTransportErrorDoesNotExposeWebhook(t *testing.T) {
	const webhookURL = "http://127.0.0.1:1/open-apis/bot/v2/hook/private-token"
	client := feishuTestClient(webhookURL, "")
	action := &ClientActionFeishu{
		Client: client,
		HTTPClient: &http.Client{
			Timeout: 100 * time.Millisecond,
		},
		Sleep: func(time.Duration) {},
	}

	err := action.pushMessage(format.PackFeishuMessage(client, "告警内容"), webhookURL, "")
	if err == nil {
		t.Fatal("transport error expected")
	}
	if strings.Contains(err.Error(), "private-token") {
		t.Fatalf("transport error exposed webhook token: %v", err)
	}
}

func TestClientActionFeishuKeepsUrlAndSecretPaired(t *testing.T) {
	client := feishuTestClient("https://example.com/first", "first-secret")
	client.ExtendFeishu.RobotUrlRandomList = append(client.ExtendFeishu.RobotUrlRandomList, "https://example.com/second")
	client.ExtendFeishu.RobotSecretRandomList = append(client.ExtendFeishu.RobotSecretRandomList, "second-secret")
	action := &ClientActionFeishu{Client: client}

	for index := 0; index < 50; index++ {
		webhookURL, secret, err := action.selectRobot()
		if err != nil {
			t.Fatalf("select robot: %v", err)
		}
		if strings.Contains(webhookURL, "first") && secret != "first-secret" {
			t.Fatalf("first webhook selected with secret %q", secret)
		}
		if strings.Contains(webhookURL, "second") && secret != "second-secret" {
			t.Fatalf("second webhook selected with secret %q", secret)
		}
	}
}
