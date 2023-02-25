package models

import (
    "errors"
    "gomessage/utils/database"
    "gorm.io/gorm"
    "strings"
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
    ID                int            `json:"id" gorm:"primarykey"`
    CreatedAt         time.Time      `json:"-"`
    UpdatedAt         time.Time      `json:"-"`
    DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
    Namespace         string         `json:"namespace"`
    ClientName        string         `json:"client_name"`
    ClientDescription string         `json:"client_description"`
    ClientType        string         `json:"client_type"`
    IsActive          bool           `json:"is_active"`
    ExtendDingtalk    *Dingtalk      `json:"-" gorm:"-:all"`
    ExtendWechat      *Wechat        `json:"-" gorm:"-:all"`
    ExtendFeishu      *Feishu        `json:"-" gorm:"-:all"`
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
        c.ExtendDingtalk.ClientId = int(c.ID)
        c.ExtendDingtalk.RobotUrl = strings.Join(c.ExtendDingtalk.RobotUrlInfoList, "\n")
        dingtalkResult := database.DB.DefaultClient.Create(&c.ExtendDingtalk)
        if dingtalkResult.Error != nil {
            return c, dingtalkResult.Error
        }

    } else if c.ClientType == "wechat" {
        c.ExtendWechat.ClientId = int(c.ID)
        wechatResult := database.DB.DefaultClient.Create(&c.ExtendWechat)
        if wechatResult.Error != nil {
            return c, wechatResult.Error
        }

    } else if c.ClientType == "feishu" {
        c.ExtendFeishu.ClientId = int(c.ID)
        c.ExtendFeishu.RobotUrl = strings.Join(c.ExtendFeishu.RobotUrls, "\n")
        feishuResult := database.DB.DefaultClient.Create(&c.ExtendFeishu)
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
        dingtalk := Dingtalk{}
        dingtalkResult := database.DB.DefaultClient.Where("client_id = ?", int(cli.ID)).First(&dingtalk)
        if dingtalkResult.Error != nil {
            return &cli, dingtalkResult.Error
        }
        dingtalk.RobotUrlInfoList = strings.Split(dingtalk.RobotUrl, "\n") //从数据库中取出机器人地址时，展开赋值给RobotUrlInfoList
        cli.ExtendDingtalk = &dingtalk

    } else if cli.ClientType == "wechat" {
        wechat := Wechat{}
        wechatResult := database.DB.DefaultClient.Where("client_id = ?", int(cli.ID)).First(&wechat)
        if wechatResult.Error != nil {
            return &cli, wechatResult.Error
        }
        cli.ExtendWechat = &wechat

    } else if cli.ClientType == "feishu" {
        feishu := Feishu{}
        feishuResult := database.DB.DefaultClient.Where("client_id = ?", int(cli.ID)).First(&feishu)
        if feishuResult.Error != nil {
            return &cli, feishuResult.Error
        }
        feishu.RobotUrls = strings.Split(feishu.RobotUrl, "\n")
        cli.ExtendFeishu = &feishu

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
    result := database.DB.DefaultClient.Where(&Client{IsActive: true}).Find(&list)
    if result.Error != nil {
        return list, result.Error
    }
    return list, nil
}
