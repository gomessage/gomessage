package models

import (
	"gorm.io/gorm"
	"time"
)

type FeishuRobot struct {
	Url    string `json:"url"`
	Secret string `json:"secret,omitempty"`
}

func JoinFeishuRobotUrl(robots []FeishuRobot) []string {
	urls := make([]string, 0, len(robots))
	for _, robot := range robots {
		urls = append(urls, robot.Url)
	}
	return urls
}

func JoinFeishuRobotSecret(robots []FeishuRobot) []string {
	secrets := make([]string, 0, len(robots))
	for _, robot := range robots {
		secrets = append(secrets, robot.Secret)
	}
	return secrets
}

type Feishu struct {
	ID                    int            `json:"id" gorm:"primarykey"`
	CreatedAt             time.Time      `json:"-"`
	UpdatedAt             time.Time      `json:"-"`
	DeletedAt             gorm.DeletedAt `json:"-" gorm:"index"`
	ClientId              int            `json:"client_id"`
	RobotKeyword          string         `json:"robot_keyword"`
	TitleColor            string         `json:"title_color"`
	RobotUrl              string         `json:"robot_url"`
	RobotSecret           string         `json:"-"`
	RobotUrlList          []FeishuRobot  `json:"robot_url_list" gorm:"-:all"`
	RobotUrlRandomList    []string       `json:"-" gorm:"-:all"`
	RobotSecretRandomList []string       `json:"-" gorm:"-:all"`
}

func (*Feishu) TableName() string {
	return "client_feishu"
}
