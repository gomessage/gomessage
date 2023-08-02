package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Url struct {
	Url string `json:"url"`
}

func JoinUrl(urls []Url) []string {
	var uList []string
	for _, v := range urls {
		uList = append(uList, v.Url)
	}
	return uList
}

type Dingtalk struct {
	ID                 int            `json:"id" gorm:"primarykey"`
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`
	ClientId           int            `json:"client_id"`
	RobotKeyword       string         `json:"robot_keyword"`
	RobotUrl           string         `json:"robot_url"`
	AtAll              bool           `json:"at_all"`
	AtMobiles          datatypes.JSON `json:"at_mobiles"`
	RobotUrlList       []Url          `gorm:"-:all" json:"robot_url_list"` //得到前端提交过来的robot_url_list字段
	RobotUrlRandomList []string       `gorm:"-:all" json:"-"`              //根据robot_url_list，得到一个url随机列表
}

func (*Dingtalk) TableName() string {
	return "client_dingtalk"
}
