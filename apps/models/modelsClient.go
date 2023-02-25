package models

import (
	"errors"
	"gomessage/utils/database"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Dingtalk struct {
	ID           int            `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
	ClientId     int            `json:"client_id"`
	RobotKeyword string         `json:"robot_keyword"`
	RobotUrl     string         `json:"robot_url"`
	RobotUrls    []string       `json:"-" gorm:"-:all"`
}

func (*Dingtalk) TableName() string {
	return "dingtalk"
}

type Feishu struct {
	ID           int            `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
	ClientId     int            `json:"client_id"`
	RobotKeyword string         `json:"robot_keyword"`
	TitleColor   string         `json:"title_color"`
	RobotUrl     string         `json:"robot_url"`
	RobotUrls    []string       `json:"-" gorm:"-:all"`
}

func (*Feishu) TableName() string {
	return "feishu"
}

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
	return "wechat"
}

type Client struct {
	ID                 int            `json:"id" gorm:"primarykey"`
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`
	Namespace          string         `json:"namespace"`
	ClientName         string         `json:"client_name"`
	ClientDescription  string         `json:"client_description"`
	ClientType         string         `json:"client_type"`
	IsActive           bool           `json:"is_active"`
	ClientInfoDingtalk *Dingtalk      `json:"-" gorm:"-:all"`
	ClientInfoWechat   *Wechat        `json:"-" gorm:"-:all"`
	ClientInfoFeishu   *Feishu        `json:"-" gorm:"-:all"`
}

func (*Client) TableName() string {
	return "client"
}

func AddClient(c *Client) (*Client, error) {
	c.IsActive = false
	createResult := database.DB.DefaultClient.Create(&c)
	if createResult.Error != nil {
		return c, createResult.Error
	}

	if c.ClientType == "dingtalk" {
		c.ClientInfoDingtalk.ClientId = int(c.ID)
		c.ClientInfoDingtalk.RobotUrl = strings.Join(c.ClientInfoDingtalk.RobotUrls, "\n")
		dingtalkResult := database.DB.DefaultClient.Create(&c.ClientInfoDingtalk)
		if dingtalkResult.Error != nil {
			return c, dingtalkResult.Error
		}

	} else if c.ClientType == "wechat" {
		c.ClientInfoWechat.ClientId = int(c.ID)
		wechatResult := database.DB.DefaultClient.Create(&c.ClientInfoWechat)
		if wechatResult.Error != nil {
			return c, wechatResult.Error
		}

	} else if c.ClientType == "feishu" {
		c.ClientInfoFeishu.ClientId = int(c.ID)
		c.ClientInfoFeishu.RobotUrl = strings.Join(c.ClientInfoFeishu.RobotUrls, "\n")
		feishuResult := database.DB.DefaultClient.Create(&c.ClientInfoFeishu)
		if feishuResult.Error != nil {
			return c, feishuResult.Error
		}

	} else {
		return c, errors.New("未知的ClientType=" + c.ClientType)
	}

	return c, nil
}

func DeleteClient(id int) (int, error) {
	var cli Client
	result := database.DB.DefaultClient.Delete(&cli, id)
	return int(result.RowsAffected), result.Error
}

func UpdateClient(id int, t *Client) (*Client, error) {
	client := Client{}
	readResult := database.DB.DefaultClient.First(&client, id)

	//如果Error不为空
	if readResult.Error != nil {
		return &client, readResult.Error

	} else {
		updateResult := database.DB.DefaultClient.Model(&client).Omit("id").Updates(&t)
		return &client, updateResult.Error
	}
}

func GetClientById(id int) (*Client, error) {
	var cli Client
	queryResult := database.DB.DefaultClient.Where("id = ?", id).First(&cli)
	if queryResult.Error != nil {
		return &cli, queryResult.Error
	}

	if cli.ClientType == "dingtalk" {
		clientInfo := Dingtalk{ClientId: int(cli.ID)}
		dingtalkResult := database.DB.DefaultClient.First(&clientInfo)
		if dingtalkResult.Error != nil {
			return &cli, dingtalkResult.Error
		}
		clientInfo.RobotUrls = strings.Split(clientInfo.RobotUrl, "\n")
		cli.ClientInfoDingtalk = &clientInfo

	} else if cli.ClientType == "wechat" {
		clientInfo := Wechat{ClientId: int(cli.ID)}
		wechatResult := database.DB.DefaultClient.First(&clientInfo)
		if wechatResult.Error != nil {
			return &cli, wechatResult.Error
		}
		cli.ClientInfoWechat = &clientInfo

	} else if cli.ClientType == "feishu" {
		clientInfo := Feishu{ClientId: int(cli.ID)}
		feishuResult := database.DB.DefaultClient.Where(&Feishu{ClientId: clientInfo.ID}).First(&clientInfo)
		if feishuResult.Error != nil {
			return &cli, feishuResult.Error
		}
		clientInfo.RobotUrls = strings.Split(clientInfo.RobotUrl, "\n")
		cli.ClientInfoFeishu = &clientInfo

	} else {
		return nil, errors.New("未知的ClientTpye=" + cli.ClientType)
	}

	return &cli, nil
}

func ListClient(ns string) (*[]Client, error) {
	var list []Client
	result := database.DB.DefaultClient.Where(&Client{Namespace: ns}).Order("id desc").Find(&list)
	if result.Error != nil {
		return &list, result.Error
	}
	return &list, nil
}

func ActiveClient(id int, active bool) (*Client, error) {
	cli := Client{}
	queryResult := database.DB.DefaultClient.First(&cli, id)
	if queryResult.Error != nil {
		return &cli, queryResult.Error

	} else {
		updateResult := database.DB.DefaultClient.Model(&cli).Updates(Client{IsActive: active})
		return &cli, updateResult.Error
	}
}

func GetActiveClient() ([]Client, error) {
	var list []Client
	result := database.DB.DefaultClient.Where(&Client{IsActive: true}).First(&list)
	if result.Error != nil {
		return list, result.Error
	}
	return list, nil
}