package models

import (
	"encoding/json"
	"errors"
	"gomessage/pkg/database"
	"gomessage/pkg/utils"
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
	ClientInfo              json.RawMessage    `gorm:"-:all" json:"client_info"`
	ExtendDingtalk          *Dingtalk          `gorm:"-:all" json:"-"`
	ExtendFeishu            *Feishu            `gorm:"-:all" json:"-"`
	ExtendWechatApplication *WechatApplication `gorm:"-:all" json:"-"`
	ExtendWechatRobot       *WechatRobot       `gorm:"-:all" json:"-"`
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
	case utils.VarDingtalk:
		c.ExtendDingtalk.ClientId = int(c.ID)
		c.ExtendDingtalk.RobotUrlRandomList = JoinUrl(c.ExtendDingtalk.RobotUrlList) //url随机列表
		c.ExtendDingtalk.RobotUrl = strings.Join(c.ExtendDingtalk.RobotUrlRandomList, "\n")
		dingtalkResult := database.DB.Default.Create(&c.ExtendDingtalk)
		if dingtalkResult.Error != nil {
			return c, dingtalkResult.Error
		}

	case utils.VarFeishu:
		c.ExtendFeishu.ClientId = int(c.ID)
		c.ExtendFeishu.RobotUrlRandomList = JoinUrl(c.ExtendFeishu.RobotUrlList) //url随机列表
		c.ExtendFeishu.RobotUrl = strings.Join(c.ExtendFeishu.RobotUrlRandomList, "\n")
		feishuResult := database.DB.Default.Create(&c.ExtendFeishu)
		if feishuResult.Error != nil {
			return c, feishuResult.Error
		}

	case utils.VarWechatRobot:
		c.ExtendWechatRobot.ClientId = int(c.ID)
		c.ExtendWechatRobot.RobotUrlRandomList = JoinUrl(c.ExtendWechatRobot.RobotUrlList) //url随机列表
		c.ExtendWechatRobot.RobotUrl = strings.Join(c.ExtendWechatRobot.RobotUrlRandomList, "\n")
		result := database.DB.Default.Create(&c.ExtendWechatRobot)
		if result.Error != nil {
			return c, result.Error
		}

	case utils.VarWechatApplication:
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

func UpdateClientInfo(id int, newClient *Client) (*Client, error) {
	oldClient := Client{}
	result := database.DB.Default.Model(&oldClient).Where("id = ? ", id).First(&oldClient)
	if result.Error != nil {
		return nil, result.Error
	}

	database.DB.Default.Model(&oldClient).Updates(Client{
		ClientName:        newClient.ClientName,
		ClientDescription: newClient.ClientDescription,
		IsActive:          newClient.IsActive,
	})

	switch newClient.ClientType {
	case utils.VarDingtalk:
		dingtalk := Dingtalk{}
		database.DB.Default.Model(&dingtalk).Where("client_id = ?", oldClient.ID).First(&dingtalk)
		newClient.ExtendDingtalk.RobotUrlRandomList = JoinUrl(newClient.ExtendDingtalk.RobotUrlList) //url随机列表
		database.DB.Default.Model(&dingtalk).Updates(Dingtalk{
			RobotKeyword: newClient.ExtendDingtalk.RobotKeyword,
			RobotUrl:     strings.Join(newClient.ExtendDingtalk.RobotUrlRandomList, "\n"),
		})

	case utils.VarFeishu:
		feishu := Feishu{}
		database.DB.Default.Model(&feishu).Where("client_id = ?", oldClient.ID).First(&feishu)
		newClient.ExtendFeishu.RobotUrlRandomList = JoinUrl(newClient.ExtendFeishu.RobotUrlList) //url随机列表
		database.DB.Default.Model(&feishu).Updates(Feishu{
			RobotKeyword: newClient.ExtendFeishu.RobotKeyword,
			TitleColor:   newClient.ExtendFeishu.TitleColor,
			RobotUrl:     strings.Join(newClient.ExtendFeishu.RobotUrlRandomList, "\n"),
		})

	case utils.VarWechatRobot:
		wechatRobot := WechatRobot{}
		database.DB.Default.Model(&wechatRobot).Where("client_id = ?", oldClient.ID).First(&wechatRobot)
		newClient.ExtendWechatRobot.RobotUrlRandomList = JoinUrl(newClient.ExtendWechatRobot.RobotUrlList) //url随机列表
		database.DB.Default.Model(&wechatRobot).Updates(WechatRobot{
			RobotKeyword: newClient.ExtendWechatRobot.RobotKeyword,
			RobotUrl:     strings.Join(newClient.ExtendWechatRobot.RobotUrlRandomList, "\n"),
		})

	case utils.VarWechatApplication:
		wechatApp := WechatApplication{}
		database.DB.Default.Model(&wechatApp).Where("client_id = ?", oldClient.ID).First(&wechatApp)
		database.DB.Default.Model(&wechatApp).Updates(WechatApplication{
			CorpId:  newClient.ExtendWechatApplication.CorpId,
			AgentId: newClient.ExtendWechatApplication.AgentId,
			Secret:  newClient.ExtendWechatApplication.Secret,
			Touser:  newClient.ExtendWechatApplication.Touser,
		})

	default:
		return newClient, errors.New("未知的ClientType=" + newClient.ClientType)
	}

	return &oldClient, result.Error

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

func GetClientById(id int) (*Client, error) {
	var cli Client
	queryResult := database.DB.Default.Where("id = ?", id).First(&cli)
	if queryResult.Error != nil {
		return &cli, queryResult.Error
	}

	switch cli.ClientType {
	case utils.VarDingtalk:
		dingtalk := Dingtalk{}
		dingtalkResult := database.DB.Default.Where("client_id = ?", int(cli.ID)).First(&dingtalk)
		if dingtalkResult.Error != nil {
			return &cli, dingtalkResult.Error
		}
		dingtalk.RobotUrlRandomList = strings.Split(dingtalk.RobotUrl, "\n") //从数据库中取出机器人地址时，展开赋值给RobotUrlInfoList
		cli.ExtendDingtalk = &dingtalk

	case utils.VarFeishu:
		feishu := Feishu{}
		feishuResult := database.DB.Default.Where("client_id = ?", int(cli.ID)).First(&feishu)
		if feishuResult.Error != nil {
			return &cli, feishuResult.Error
		}
		feishu.RobotUrlRandomList = strings.Split(feishu.RobotUrl, "\n")
		cli.ExtendFeishu = &feishu

	case utils.VarWechatRobot:
		wechatRobot := WechatRobot{}
		wechatRobotResult := database.DB.Default.Where("client_id = ?", int(cli.ID)).First(&wechatRobot)
		if wechatRobotResult.Error != nil {
			return &cli, wechatRobotResult.Error
		}
		wechatRobot.RobotUrlRandomList = strings.Split(wechatRobot.RobotUrl, "\n")
		cli.ExtendWechatRobot = &wechatRobot

	case utils.VarWechatApplication:
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
