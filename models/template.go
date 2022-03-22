package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Templates struct {
	Id              int       `json:"id" orm:"pk;auto"`                    //id，主键，自增
	NamespaceId     int       `json:"namespace_id"`                        //外键连接Namespace表
	TemplateName    string    `json:"template_name" orm:"size(500)"`       //消息模板的标签或名称
	MessageTemplate string    `json:"message_template" orm:"size(2000)"`   //消息模板的主体（也就是变量与字符串的混合体）
	MessageMerge    bool      `json:"message_merge"`                       //消息模板是否切割渲染
	CreateTime      time.Time `json:"-" orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateTime      time.Time `json:"-" orm:"auto_now;type(datetime)"`     //更新时间
}

func init() {
	orm.RegisterModel(new(Templates))
}

//######################
//函数：查询和创建一条数据，确保它始终存在
//######################
func ReadOrCreateTemplate(templateName string, messageTemplate string, messageMerge bool, ns *Namespaces) Templates {

	if ns == nil {
		fmt.Println("找不到所属Namespace，不进行Template创建...")
		return Templates{}
	} else {
		o := orm.NewOrm()
		t2 := Templates{
			TemplateName:    templateName,
			MessageTemplate: messageTemplate,
			MessageMerge:    messageMerge,
			NamespaceId:     ns.Id,
		}
		newCreate, id, err := o.ReadOrCreate(&t2, "TemplateName", "NamespaceId")
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
//函数：删除用户
//######################
func DeleteTemplate(temp Templates) int64 {
	o := orm.NewOrm()
	num, err2 := o.Delete(&temp)
	if err2 != nil {
		fmt.Println(err2)
	}
	return num
}

//######################
//函数：查询一条数据
//######################
func GetOneTemplateId(id int) Templates {
	var temp Templates
	o := orm.NewOrm()
	err := o.QueryTable("templates").Filter("id", id).One(&temp)
	if err == orm.ErrMultiRows {
		fmt.Printf("报错：找到的数据不是一条")
	}
	if err == orm.ErrNoRows {
		fmt.Printf("报错：没有查询到Template数据")
	}
	return temp
}

//######################
//函数：查询一条数据
//######################
func GetOneTemplateNsId(id int) Templates {
	var temp Templates
	o := orm.NewOrm()
	err := o.QueryTable("templates").Filter("namespace_id", id).One(&temp)
	if err == orm.ErrMultiRows {
		fmt.Printf("报错：找到的数据不是一条")
	}
	if err == orm.ErrNoRows {
		fmt.Printf("报错：没有查询到Template数据")
	}
	return temp
}

//######################
//函数：查询所有（按namespace查询）
//######################
func ListNsTemplate(ns *Namespaces) []Templates {
	var TemplateList []Templates
	o := orm.NewOrm()
	num, err := o.QueryTable("templates").Filter("namespace_id", &ns.Id).All(&TemplateList)
	if err != nil {
		return nil
	}
	fmt.Printf("查询全部ConfigMap，受影响的行数：%v\n", num)
	return TemplateList
}

//######################
//函数：查询全部数据（不附加任何条件）
//######################
func ListAllTemplate() ([]*Templates, error) {
	var list []*Templates
	o := orm.NewOrm()
	_, err := o.QueryTable(Templates{}).OrderBy("-id").All(&list)
	if err != nil {
		return nil, err
	}
	return list, err
}
