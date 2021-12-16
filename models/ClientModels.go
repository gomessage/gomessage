package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type DingtalkClient struct {
	Id           int       `orm:"pk;auto"`
	RobotName    string    `orm:"size(500)"`
	RobotUrl     string    `orm:"size(2000)"`
	RobotKeyword string    `orm:"size(2000)"`
	CreateTime   time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
}

func init() {
	orm.RegisterModel(new(DingtalkClient))
}

//######################
//增
//######################
func AddOneDingtalkClient(m DingtalkClient) int64 {
	o := orm.NewOrm()
	uid, err := o.Insert(&m)
	if err != nil {
		fmt.Errorf("数据插入失败")
	}
	return uid
}

//######################
//删
//######################
func DeleteOneDingtalkClient(m DingtalkClient) int64 {
	o := orm.NewOrm()
	num, err2 := o.Delete(&m)
	if err2 != nil {
		fmt.Println(err2)
	}
	return num
}

//######################
//查：查询一条数据
//######################
func QueryOneDingtalkClient(value string) DingtalkClient {
	var m DingtalkClient
	o := orm.NewOrm()
	err := o.QueryTable(&m).Filter("Label", value).One(&m)
	if err == orm.ErrMultiRows {
		fmt.Printf("报错：找到的数据不是一条")
	}
	if err == orm.ErrNoRows {
		fmt.Printf("报错：没有查询到对应数据")
	}
	return m
}

//######################
//查：查询所有数据
//######################
func QueryAllDingtalkClient() []DingtalkClient {
	var mList []DingtalkClient
	m := DingtalkClient{}
	o := orm.NewOrm()
	num, _ := o.QueryTable(&m).All(&mList)
	fmt.Printf("受影响的行数：%v\n", num)
	return mList
}
