package models

import (
	"errors"
	"gomessage/apps/models/clients"
	"gomessage/utils/database"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Client struct {
	ID                      int                        `json:"id" gorm:"primarykey"`
	CreatedAt               time.Time                  `json:"-"`
	UpdatedAt               time.Time                  `json:"-"`
	DeletedAt               gorm.DeletedAt             `json:"-" gorm:"index"`
	Namespace               string                     `json:"namespace"`
	ClientName              string                     `json:"client_name"`
	ClientDescription       string                     `json:"client_description"`
	ClientType              string                     `json:"client_type"`
	IsActive                bool                       `json:"is_active"`
	ExtendDingtalk          *clients.Dingtalk          `json:"-" gorm:"-:all"`
	ExtendFeishu            *clients.Feishu            `json:"-" gorm:"-:all"`
	ExtendWechatApplication *clients.WechatApplication `json:"-" gorm:"-:all"`
	ExtendWechatRobot       *clients.WechatRobot       `json:"-" gorm:"-:all"`
}

func (*Client) TableName() string {
	return "clients"
}

func AddClient(c *Client) (*Client, error) {
	c.IsActive = false
	createResult := database.DB.DefaultClient.Create(&c)
	if createResult.Error != nil {
		return c, createResult.Error
	}

	switch c.ClientType {
	case "dingtalk":
		c.ExtendDingtalk.ClientId = int(c.ID)
		c.ExtendDingtalk.RobotUrl = strings.Join(c.ExtendDingtalk.RobotUrlInfoList, "\n")
		dingtalkResult := database.DB.DefaultClient.Create(&c.ExtendDingtalk)
		if dingtalkResult.Error != nil {
			return c, dingtalkResult.Error
		}

	case "feishu":
		c.ExtendFeishu.ClientId = int(c.ID)
		c.ExtendFeishu.RobotUrl = strings.Join(c.ExtendFeishu.RobotUrlInfoList, "\n")
		feishuResult := database.DB.DefaultClient.Create(&c.ExtendFeishu)
		if feishuResult.Error != nil {
			return c, feishuResult.Error
		}

	case "wechat_robot":
		c.ExtendWechatRobot.ClientId = int(c.ID)
		c.ExtendWechatRobot.RobotUrl = strings.Join(c.ExtendWechatRobot.RobotUrlInfoList, "\n")
		result := database.DB.DefaultClient.Create(&c.ExtendWechatRobot)
		if result.Error != nil {
			return c, result.Error
		}

	case "wechat":
		c.ExtendWechatApplication.ClientId = int(c.ID)
		wechatResult := database.DB.DefaultClient.Create(&c.ExtendWechatApplication)
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
	result := database.DB.DefaultClient.Delete(&cli, id)
	return int(result.RowsAffected), result.Error
}

func UpdateClientActive(id int, t *Client) (*Client, error) {
	client := Client{}
	updateResult := database.DB.DefaultClient.Model(&client).Where("id = ? ", id).Update("is_active", t.IsActive)
	return &client, updateResult.Error

}

func UpdateClientInfo(id int, t *Client) (*Client, error) {
	client := Client{}
	updateResult := database.DB.DefaultClient.Model(&client).Where("id = ? ", id).Updates(map[string]any{"is_active": t.IsActive, "client_name": t.ClientName, "client_description": t.ClientDescription})
	return &client, updateResult.Error

}

func GetClientById(id int) (*Client, error) {
	var cli Client
	queryResult := database.DB.DefaultClient.Where("id = ?", id).First(&cli)
	if queryResult.Error != nil {
		return &cli, queryResult.Error
	}

	switch cli.ClientType {
	case "dingtalk":
		dingtalk := clients.Dingtalk{}
		dingtalkResult := database.DB.DefaultClient.Where("client_id = ?", int(cli.ID)).First(&dingtalk)
		if dingtalkResult.Error != nil {
			return &cli, dingtalkResult.Error
		}
		dingtalk.RobotUrlInfoList = strings.Split(dingtalk.RobotUrl, "\n") //从数据库中取出机器人地址时，展开赋值给RobotUrlInfoList
		cli.ExtendDingtalk = &dingtalk

	case "feishu":
		feishu := clients.Feishu{}
		feishuResult := database.DB.DefaultClient.Where("client_id = ?", int(cli.ID)).First(&feishu)
		if feishuResult.Error != nil {
			return &cli, feishuResult.Error
		}
		feishu.RobotUrlInfoList = strings.Split(feishu.RobotUrl, "\n")
		cli.ExtendFeishu = &feishu

	case "wechat_robot":
		wechatRobot := clients.WechatRobot{}
		wechatRobotResult := database.DB.DefaultClient.Where("client_id = ?", int(cli.ID)).First(&wechatRobot)
		if wechatRobotResult.Error != nil {
			return &cli, wechatRobotResult.Error
		}
		wechatRobot.RobotUrlInfoList = strings.Split(wechatRobot.RobotUrl, "\n")
		cli.ExtendWechatRobot = &wechatRobot

	case "wechat":
		wechat := clients.WechatApplication{}
		wechatResult := database.DB.DefaultClient.Where("client_id = ?", int(cli.ID)).First(&wechat)
		if wechatResult.Error != nil {
			return &cli, wechatResult.Error
		}
		cli.ExtendWechatApplication = &wechat

	default:
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

// GetActiveClient 获取指定命名空间下的处于激活状态的客户端
func GetActiveClient(ns string) ([]Client, error) {
	var list []Client
	result := database.DB.DefaultClient.Where(&Client{IsActive: true, Namespace: ns}).Find(&list)
	if result.Error != nil {
		return list, result.Error
	}
	return list, nil
}
