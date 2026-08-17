package models

import (
	"path/filepath"
	"testing"

	"github.com/spf13/viper"

	"gomessage/pkg/utils"
	"gomessage/pkg/utils/database"
)

func prepareFeishuTestDatabase(t *testing.T) {
	t.Helper()
	viper.Set("sqlite3.path", filepath.Join(t.TempDir(), "gomessage.db"))
	viper.Set("sqlite3.MaxIdleConns", 1)
	viper.Set("sqlite3.MaxOpenConns", 1)
	database.DB.Init("sqlite")
	if err := database.DB.Default.AutoMigrate(&Client{}, &Feishu{}); err != nil {
		t.Fatalf("migrate test database: %v", err)
	}
}

func TestFeishuRobotSecretsRoundTrip(t *testing.T) {
	prepareFeishuTestDatabase(t)

	client := &Client{
		Namespace:  "test",
		ClientName: "feishu",
		ClientType: utils.VarFeishu,
		ExtendFeishu: &Feishu{
			RobotKeyword: "告警",
			TitleColor:   "red",
			RobotUrlList: []FeishuRobot{
				{Url: "https://example.com/first", Secret: "first-secret"},
				{Url: "https://example.com/second", Secret: "second-secret"},
			},
		},
	}
	if _, err := AddClient(client); err != nil {
		t.Fatalf("add client: %v", err)
	}

	stored, err := GetClientById(client.ID)
	if err != nil {
		t.Fatalf("get client: %v", err)
	}
	if got, want := stored.ExtendFeishu.RobotSecretRandomList, []string{"first-secret", "second-secret"}; !equalStrings(got, want) {
		t.Fatalf("stored secrets = %#v, want %#v", got, want)
	}

	client.ExtendFeishu.RobotUrlList = []FeishuRobot{{Url: "https://example.com/first"}}
	if _, err := UpdateClientInfo(client.ID, client); err != nil {
		t.Fatalf("update client: %v", err)
	}
	stored, err = GetClientById(client.ID)
	if err != nil {
		t.Fatalf("get updated client: %v", err)
	}
	if got, want := stored.ExtendFeishu.RobotSecretRandomList, []string{""}; !equalStrings(got, want) {
		t.Fatalf("cleared secrets = %#v, want %#v", got, want)
	}
}

func equalStrings(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
