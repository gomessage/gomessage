package client

import (
	"encoding/json"
	"strings"
	"testing"

	"gomessage/pkg/models"
)

func TestRequestDataFeishuExposesPerRobotSecretOnly(t *testing.T) {
	payload, err := json.Marshal(RequestDataFeishu{
		Feishu: &models.Feishu{RobotSecret: "stored-secret"},
		RobotUrlList: []models.FeishuRobot{{
			Url:    "https://example.com/hook",
			Secret: "robot-secret",
		}},
	})
	if err != nil {
		t.Fatalf("marshal Feishu client response: %v", err)
	}

	response := string(payload)
	if !strings.Contains(response, `"secret":"robot-secret"`) {
		t.Fatalf("per-robot secret missing from response: %s", response)
	}
	if strings.Contains(response, "stored-secret") || strings.Contains(response, "robot_secret") {
		t.Fatalf("internal secret storage leaked in response: %s", response)
	}
}
