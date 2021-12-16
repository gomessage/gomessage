package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Template struct {
	Id              int       `orm:"pk;auto"`
	Label           string    `orm:"size(500)"`
	MessageTemplate string    `orm:"size(2000)"`
	CreateTime      time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
}

func init() {
	orm.RegisterModel(new(Template))
}

//######################
//函数：查询和创建一条数据，确保它始终存在
//######################
func ReadOrCreateTemplate(label string, messageTemplate string) Template {
	o := orm.NewOrm()
	t2 := Template{}
	t2.Label = label
	t2.MessageTemplate = messageTemplate
	newCreate, id, err := o.ReadOrCreate(&t2, "Label")

	if err != nil {
		panic(err)
	}

	//判断用户是否是新创建的
	if newCreate {
		fmt.Println("新创建用户ID：", id)
	} else {
		fmt.Println("查询到用户ID：", id)
	}

	return t2
}

//######################
//函数：查询一条数据
//######################
func GetOneTemplate(lable string) Template {
	var temp Template
	o := orm.NewOrm()
	err := o.QueryTable(&temp).Filter("Label", lable).One(&temp)
	if err == orm.ErrMultiRows {
		fmt.Printf("报错：找到的数据不是一条")
	}
	if err == orm.ErrNoRows {
		fmt.Printf("报错：没有查询到对应数据")
	}
	return temp
}

//######################
//函数：删除用户
//######################
func DeleteTemplate(j Template) int64 {
	o := orm.NewOrm()
	num, err2 := o.Delete(&j)
	if err2 != nil {
		fmt.Println(err2)
	}
	return num
}

//######################
//函数：查询所有用户
//######################
func GetAllTemplate() []Template {
	var temp []Template
	t := Template{}

	o := orm.NewOrm()
	num, _ := o.QueryTable(&t).All(&temp)
	fmt.Printf("受影响的行数：%v\n", num)

	return temp
}
