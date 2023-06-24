package clients

import (
	"gorm.io/gorm"
	"time"
)

type Feishu struct {
	ID               int            `json:"id" gorm:"primarykey"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
	ClientId         int            `json:"client_id"`
	RobotKeyword     string         `json:"robot_keyword"`
	TitleColor       string         `json:"title_color"`
	RobotUrl         string         `json:"robot_url"`
	RobotUrlInfoList []string       `json:"-" gorm:"-:all"`
}

func (*Feishu) TableName() string {
	return "client_feishu"
}
