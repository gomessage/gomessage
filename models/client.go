package models

import (
	"encoding/json"
	"errors"
	"gomessage/pkg/database"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Client struct {
	ID                      int                `json:"id" gorm:"primarykey"`
	CreatedAt               time.Time          `json:"-"`
	UpdatedAt               time.Time          `json:"-"`
	DeletedAt               gorm.DeletedAt     `json:"-" gorm:"index"`
	Namespace               string             `json:"namespace"`
	ClientName              string             `json:"client_name"`
	ClientDescription       string             `json:"client_description"`
	ClientType              string             `json:"client_type"`
	IsActive                bool               `json:"is_active"`
	ClientInfo              json.RawMessage    `json:"client_info" gorm:"-:all"`
	ExtendDingtalk          *Dingtalk          `json:"-" gorm:"-:all"`
	ExtendFeishu            *Feishu            `json:"-" gorm:"-:all"`
	ExtendWechatApplication *WechatApplication `json:"-" gorm:"-:all"`
	ExtendWechatRobot       *WechatRobot       `json:"-" gorm:"-:all"`
}

func (*Client) TableName() string {
	return "clients"
}

func AddClient(c *Client) (*Client, error) {
	c.IsActive = false
	createResult := database.DB.Default.Create(&c)
	if createResult.Error != nil {
		return c, createResult.Error
	}

	switch c.ClientType {
	case "dingtalk":
		c.ExtendDingtalk.ClientId = int(c.ID)
		c.ExtendDingtalk.RobotUrl = strings.Join(c.ExtendDingtalk.RobotUrlRandomList, "\n")
		dingtalkResult := database.DB.Default.Create(&c.ExtendDingtalk)
		if dingtalkResult.Error != nil {
			return c, dingtalkResult.Error
		}

	case "feishu":
		c.ExtendFeishu.ClientId = int(c.ID)
		c.ExtendFeishu.RobotUrl = strings.Join(c.ExtendFeishu.RobotUrlRandomList, "\n")
		feishuResult := database.DB.Default.Create(&c.ExtendFeishu)
		if feishuResult.Error != nil {
			return c, feishuResult.Error
		}

	case "wechat_robot":
		c.ExtendWechatRobot.ClientId = int(c.ID)
		c.ExtendWechatRobot.RobotUrl = strings.Join(c.ExtendWechatRobot.RobotUrlRandomList, "\n")
		result := database.DB.Default.Create(&c.ExtendWechatRobot)
		if result.Error != nil {
			return c, result.Error
		}

	case "wechat":
		c.ExtendWechatApplication.ClientId = int(c.ID)
		wechatResult := database.DB.Default.Create(&c.ExtendWechatApplication)
		if wechatResult.Error != nil {
			return c, wechatResult.Error
		}

	default:
		return c, errors.New("未知的ClientType=" + c.ClientType)
	}

	return c, nil
}

func DeleteClient(id int) (int, error) {
	var cli Client
	result := database.DB.Default.Delete(&cli, id)
	return int(result.RowsAffected), result.Error
}

func UpdateClientActive(id int, t *Client) (*Client, error) {
	client := Client{}
	updateResult := database.DB.Default.Model(&client).Where("id = ? ", id).Update("is_active", t.IsActive)
	return &client, updateResult.Error

}

func UpdateClientInfo(id int, t *Client) (*Client, error) {
	client := Client{}
	updateResult := database.DB.Default.Model(&client).Where("id = ? ", id).Updates(map[string]any{"is_active": t.IsActive, "client_name": t.ClientName, "client_description": t.ClientDescription})
	return &client, updateResult.Error

}

func GetClientById(id int) (*Client, error) {
	var cli Client
	queryResult := database.DB.Default.Where("id = ?", id).First(&cli)
	if queryResult.Error != nil {
		return &cli, queryResult.Error
	}

	switch cli.ClientType {
	case "dingtalk":
		dingtalk := Dingtalk{}
		dingtalkResult := database.DB.Default.Where("client_id = ?", int(cli.ID)).First(&dingtalk)
		if dingtalkResult.Error != nil {
			return &cli, dingtalkResult.Error
		}
		dingtalk.RobotUrlRandomList = strings.Split(dingtalk.RobotUrl, "\n") //从数据库中取出机器人地址时，展开赋值给RobotUrlInfoList
		cli.ExtendDingtalk = &dingtalk

	case "feishu":
		feishu := Feishu{}
		feishuResult := database.DB.Default.Where("client_id = ?", int(cli.ID)).First(&feishu)
		if feishuResult.Error != nil {
			return &cli, feishuResult.Error
		}
		feishu.RobotUrlRandomList = strings.Split(feishu.RobotUrl, "\n")
		cli.ExtendFeishu = &feishu

	case "wechat_robot":
		wechatRobot := WechatRobot{}
		wechatRobotResult := database.DB.Default.Where("client_id = ?", int(cli.ID)).First(&wechatRobot)
		if wechatRobotResult.Error != nil {
			return &cli, wechatRobotResult.Error
		}
		wechatRobot.RobotUrlRandomList = strings.Split(wechatRobot.RobotUrl, "\n")
		cli.ExtendWechatRobot = &wechatRobot

	case "wechat":
		wechatApplication := WechatApplication{}
		wechatApplicationResult := database.DB.Default.Where("client_id = ?", int(cli.ID)).First(&wechatApplication)
		if wechatApplicationResult.Error != nil {
			return &cli, wechatApplicationResult.Error
		}
		cli.ExtendWechatApplication = &wechatApplication

	default:
		return nil, errors.New("未知的ClientType=" + cli.ClientType)
	}

	return &cli, nil
}

func ListClient(ns string) (*[]Client, error) {
	var list []Client
	result := database.DB.Default.Where(&Client{Namespace: ns}).Order("id desc").Find(&list)
	if result.Error != nil {
		return &list, result.Error
	}
	return &list, nil
}

// GetActiveClient 获取指定命名空间下的处于激活状态的客户端
func GetActiveClient(ns string) ([]Client, error) {
	var list []Client
	result := database.DB.Default.Where(&Client{IsActive: true, Namespace: ns}).Find(&list)
	if result.Error != nil {
		return list, result.Error
	}
	return list, nil
}
