package models

import (
	"errors"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Client struct {
	Id          int       `orm:"pk;auto" json:"id"`
	Name        string    `orm:"size(50)" json:"client_name"`
	Description string    `orm:"size(200)" json:"client_description"`
	Active      bool      `json:"client_active"`
	Type        string    `json:"client_type"`
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)" json:"-"`

	ClientDingtalk *ClientDingtalk `orm:"-" json:"-"`
	ClientWechat   *ClientWechat   `orm:"-" json:"-"`
	ClientFeishu   *ClientFeishu   `orm:"-" json:"-"`
}

type ClientDingtalk struct {
	ClientId     int       `orm:"pk" json:"-"`
	RobotUrl     string    `orm:"size(5000)" json:"-"`
	RobotUrlList []string  `orm:"-" json:"robot_url"`
	RobotKeyword string    `orm:"size(50)" json:"robot_keyword"`
	CreateTime   time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	UpdateTime   time.Time `orm:"auto_now;type(datetime)" json:"-"`
}

type ClientWechat struct {
	ClientId   int       `orm:"pk" json:"-"`
	Corpid     string    `orm:"size(500)" json:"wechat_corpid"`
	Agentid    string    `orm:"size(500)" json:"wechat_agentid"`
	Secret     string    `orm:"size(500)" json:"wechat_secret"`
	Touser     string    `orm:"size(500)" json:"wechat_touser"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)" json:"-"`
}

type ClientFeishu struct {
	ClientId     int       `orm:"pk" json:"-"`
	RobotUrl     string    `orm:"size(5000)" json:"-"`
	RobotUrlList []string  `orm:"-" json:"robot_url"`
	RobotKeyword string    `orm:"size(50)" json:"robot_keyword"`
	TitleColor   string    `orm:"size(50)" json:"title_color"`
	CreateTime   time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	UpdateTime   time.Time `orm:"auto_now;type(datetime)" json:"-"`
}

func init() {
	orm.RegisterModel(
		new(Client),
		new(ClientDingtalk),
		new(ClientWechat),
		new(ClientFeishu),
	)
}

func ListClient() ([]*Client, error) {
	list := []*Client{}
	o := orm.NewOrm()
	_, err := o.QueryTable(&Client{}).OrderBy("-id").All(&list)
	return list, err
}

func GetClient(id int) (*Client, error) {
	o := orm.NewOrm()

	one := Client{Id: id}
	err := o.Read(&one)
	if err != nil {
		return nil, err
	}

	if one.Type == "dingtalk" {
		clientDingtalk := ClientDingtalk{ClientId: one.Id}
		err = o.Read(&clientDingtalk)
		if err != nil {
			return nil, err
		}
		clientDingtalk.RobotUrlList = strings.Split(clientDingtalk.RobotUrl, "\n")
		one.ClientDingtalk = &clientDingtalk

	} else if one.Type == "wechat" {
		clientWechat := ClientWechat{ClientId: one.Id}
		err = o.Read(&clientWechat)
		if err != nil {
			return nil, err
		}
		one.ClientWechat = &clientWechat

	} else if one.Type == "feishu" {
		clientFeishu := ClientFeishu{ClientId: one.Id}
		err = o.Read(&clientFeishu)
		if err != nil {
			return nil, err
		}
		clientFeishu.RobotUrlList = strings.Split(clientFeishu.RobotUrl, "\n")
		one.ClientFeishu = &clientFeishu

	} else {
		return nil, errors.New("未知的client_type=" + one.Type)
	}

	return &one, nil
}

func AddClient(c *Client) (int64, error) {
	o := orm.NewOrm()

	c.Active = false

	var n1 int64
	var err error
	n1, err = o.Insert(c)
	if err != nil {
		return n1, err
	}

	var n2 int64
	if c.Type == "dingtalk" {
		c.ClientDingtalk.ClientId = c.Id
		c.ClientDingtalk.RobotUrl = strings.Join(c.ClientDingtalk.RobotUrlList, "\n")
		n2, err = o.Insert(c.ClientDingtalk)
		if err != nil {
			return n1 + n2, err
		}

	} else if c.Type == "wechat" {
		c.ClientWechat.ClientId = c.Id
		n2, err = o.Insert(c.ClientWechat)
		if err != nil {
			return n1 + n2, err
		}

	} else if c.Type == "feishu" {
		c.ClientFeishu.ClientId = c.Id
		c.ClientFeishu.RobotUrl = strings.Join(c.ClientFeishu.RobotUrlList, "\n")
		n2, err = o.Insert(c.ClientFeishu)
		if err != nil {
			return n1 + n2, err
		}

	} else {
		return 0, errors.New("未知的client_type=" + c.Type)
	}

	return n1 + n2, nil
}

func DelClient(id int) (int64, error) {
	o := orm.NewOrm()

	one := Client{Id: id}
	err := o.Read(&one)
	if err != nil {
		return 0, err
	}

	var n1 int64
	if one.Type == "dingtalk" {
		n1, err = o.Delete(&ClientDingtalk{ClientId: one.Id})
		if err != nil {
			return n1, err
		}
	} else if one.Type == "wechat" {
		n1, err = o.Delete(&ClientWechat{ClientId: one.Id})
		if err != nil {
			return n1, err
		}
	} else if one.Type == "feishu" {
		n1, err = o.Delete(&ClientFeishu{ClientId: one.Id})
		if err != nil {
			return n1, err
		}
	}

	var n2 int64
	n2, err = o.Delete(&one)
	if err != nil {
		return n1 + n2, err
	}
	return n1 + n2, nil
}

func ActiveClient(id int, active bool) (int64, error) {
	o := orm.NewOrm()
	n, err := o.Update(&Client{Id: id, Active: active}, "Active")
	return n, err
}

func GetClentActive() []Client {
	var list []Client
	o := orm.NewOrm()
	_, err := o.QueryTable(&Client{}).Filter("active", "1").All(&list)
	if err != nil {
		return nil
	}
	//fmt.Println(all)
	return list
}
