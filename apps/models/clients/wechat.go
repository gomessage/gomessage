package clients

import (
	"gorm.io/gorm"
	"time"
)

type Wechat struct {
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

func (*Wechat) TableName() string {
	return "client_wechat"
}
