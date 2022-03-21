package models

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"strings"
	"time"
)

type Client struct {
	Id             int             `json:"id" orm:"pk;auto"`                    //主键，自增
	NamespaceId    int             `json:"namespace_id"`                        //Namespace表中的主键ID，表结构上没有做外键，仅仅是在代码层面完成关系绑定
	Name           string          `json:"client_name" orm:"size(50)"`          //客户端名称
	Description    string          `json:"client_description" orm:"size(200)"`  //客户端描述
	Active         bool            `json:"client_active"`                       //客户端是否被激活
	Type           string          `json:"client_type" orm:"size(50)"`          //客户端类型
	CreateTime     time.Time       `json:"-" orm:"auto_now_add;type(datetime)"` //添加时间
	UpdateTime     time.Time       `json:"-" orm:"auto_now;type(datetime)"`     //修改时间
	ClientDingtalk *ClientDingtalk `json:"-" orm:"-"`                           //钉钉客户端，结构体中包含，但是orm中不创建
	ClientWechat   *ClientWechat   `json:"-" orm:"-"`                           //企业微信客户端，结构体中包含，但是orm中不创建
	ClientFeishu   *ClientFeishu   `json:"-" orm:"-"`                           //飞书客户端，结构体中包含，但是orm中不创建
}

type ClientDingtalk struct {
	ClientId     int       `json:"-" orm:"pk"`                          //主键，不自增
	RobotUrl     string    `json:"-" orm:"size(5000)"`                  //机器人URL
	RobotUrlList []string  `json:"robot_url" orm:"-"`                   //机器人List，结构体中包含，但是orm中不创建
	RobotKeyword string    `json:"robot_keyword" orm:"size(50)"`        //机器人放行关键字
	CreateTime   time.Time `json:"-" orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateTime   time.Time `json:"-" orm:"auto_now;type(datetime)"`     //更新时间
}

type ClientWechat struct {
	ClientId   int       `json:"-" orm:"pk"`                          //主键，不自增
	Corpid     string    `json:"wechat_corpid" orm:"size(500)"`       //企业ID
	Agentid    string    `json:"wechat_agentid" orm:"size(500)"`      //agentID
	Secret     string    `json:"wechat_secret" orm:"size(500)"`       //秘钥
	Touser     string    `json:"wechat_touser" orm:"size(500)"`       //收件人用户ID
	CreateTime time.Time `json:"-" orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateTime time.Time `json:"-" orm:"auto_now;type(datetime)"`     //更新时间
}

type ClientFeishu struct {
	ClientId     int       `json:"-" orm:"pk"`                          //主键、不自增
	RobotUrl     string    `json:"-" orm:"size(5000)"`                  //机器人URL
	RobotUrlList []string  `json:"robot_url" orm:"-"`                   //机器人List，结构体中包含，但是orm中不创建
	RobotKeyword string    `json:"robot_keyword" orm:"size(50)"`        //机器人放行关键字
	TitleColor   string    `json:"title_color" orm:"size(50)"`          //标题颜色
	CreateTime   time.Time `json:"-" orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateTime   time.Time `json:"-" orm:"auto_now;type(datetime)"`     //更新时间
}

func init() {
	orm.RegisterModel(
		new(Client),
		new(ClientDingtalk),
		new(ClientWechat),
		new(ClientFeishu),
	)
}

//查询所有客户端
func ListClient() ([]*Client, error) {
	list := []*Client{}
	o := orm.NewOrm()
	_, err := o.QueryTable(&Client{}).OrderBy("-id").All(&list)
	return list, err
}

//查询指定namespace下的客户端
func ListNsClient(ns *Namespaces) ([]*Client, error) {
	var list []*Client
	o := orm.NewOrm()
	_, err := o.QueryTable(&Client{}).Filter("namespace_id", ns.Id).OrderBy("-id").All(&list)
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
