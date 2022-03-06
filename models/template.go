package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Templates struct {
	Id              int         `orm:"pk;auto" json:"id"`                    //id，主键，自增
	Namespace       *Namespaces `orm:"rel(fk);cascade"`                      //外键连接Namespace表
	Label           string      `orm:"size(500)" json:"label"`               //消息模板的标签或名称
	MessageTemplate string      `orm:"size(2000)" json:"message_template"`   //消息模板的主体（也就是变量与字符串的混合体）
	MessageMerge    bool        `json:"message_merge"`                       //消息模板是否切割渲染
	CreateTime      time.Time   `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateTime      time.Time   `orm:"auto_now;type(datetime)" json:"-"`     //更新时间
}

func init() {
	orm.RegisterModel(new(Templates))
}

//######################
//函数：查询和创建一条数据，确保它始终存在
//######################
func ReadOrCreateTemplate(label string, messageTemplate string, MessageMerge bool, namespace *Namespaces) Templates {

	if namespace == nil {
		fmt.Println("找不到所属Namespace，不进行Template创建...")
		return Templates{}
	} else {
		o := orm.NewOrm()
		t2 := Templates{
			Label:           label,
			MessageTemplate: messageTemplate,
			MessageMerge:    MessageMerge,
			Namespace:       namespace,
		}
		newCreate, id, err := o.ReadOrCreate(&t2, "Label")
		if err != nil {
			panic(err)
		}
		//判断是否是新创建的
		if newCreate {
			fmt.Println("新创建template的ID：", id)
		} else {
			fmt.Println("查询到template的ID：", id)
		}
		return t2
	}

}

//######################
//函数：查询一条数据
//######################
func GetOneTemplate(lable string) Templates {
	var temp Templates
	o := orm.NewOrm()
	err := o.QueryTable(&temp).Filter("Label", lable).One(&temp)
	if err == orm.ErrMultiRows {
		fmt.Printf("报错：找到的数据不是一条")
	}
	if err == orm.ErrNoRows {
		fmt.Printf("报错：没有查询到Template数据")
	}
	return temp
}

//######################
//函数：删除用户
//######################
func DeleteTemplate(j Templates) int64 {
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
func QueryAllTemplate() []Templates {
	var temp []Templates
	t := Templates{}

	o := orm.NewOrm()
	_, _ = o.QueryTable(&t).All(&temp)
	//fmt.Printf("查询全部template，受影响的行数：%v\n", num)

	return temp
}
