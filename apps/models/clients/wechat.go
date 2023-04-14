package clients

import (
	"gorm.io/gorm"
	"time"
)

// WechatApplication 企业微信-应用号
type WechatApplication struct {
	ID        int            `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	ClientId  int            `json:"client_id"`
	CorpId    string         `json:"corp_id"`
	AgentId   string         `json:"agent_id"`
	Secret    string         `json:"secret"`
	Touser    string         `json:"touser"`
}

func (*WechatApplication) TableName() string {
	return "client_wechat_application"
}

// WechatRobot 企业微信-机器人
type WechatRobot struct {
	ID            int            `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
	ClientId      int            `json:"client_id"`
	RobotKeyword  string         `json:"robot_keyword"`
	RobotUrl      string         `json:"robot_url"`
	RobotUrlsList []string       `json:"-" gorm:"-:all"`
}

func (*WechatRobot) TableName() string {
	return "client_wechat_robot"
}
