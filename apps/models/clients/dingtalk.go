package clients

import (
	"gorm.io/gorm"
	"time"
)

type Dingtalk struct {
	ID               int            `json:"id" gorm:"primarykey"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
	ClientId         int            `json:"client_id"`
	RobotKeyword     string         `json:"robot_keyword"`
	RobotUrl         string         `json:"robot_url"`
	RobotUrlInfoList []string       `json:"-" gorm:"-:all"`
}

func (*Dingtalk) TableName() string {
	return "client_dingtalk"
}
