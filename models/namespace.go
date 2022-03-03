package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(Namespace))
}

//命名空间（表结构）
type Namespace struct {
	Id          int       `orm:"pk;auto"`                     //id
	Name        string    `orm:"size(50)"`                    //命名空间
	Description string    `orm:"size(200)"`                   //描述
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`     //修改时间
}

//增
func AddNamespace(n *Namespace) (int64, error) {
	o := orm.NewOrm()

	var n1 int64
	var err error
	n1, err = o.Insert(n)
	if err != nil {
		return 0, err
	}
	return n1, err
}

//删
func DelNamespace(id int) (int64, error) {
	o := orm.NewOrm()
	one := Namespace{Id: id}
	n1, err := o.Delete(&one)
	if err != nil {
		return 0, err
	}
	return n1, err
}

//改：（这个改方法可能写的不正确，后面有空了再修改）
func UpdateNamespace(id int, namespace Namespace) (int64, error) {
	o := orm.NewOrm()
	one := Namespace{Id: id}

	var n1 int64
	var err error

	if o.Read(&one) == nil {
		one = namespace
		n1, err = o.Update(&one)
		if err != nil {
			return 0, err
		}
		return n1, err
	}
	return 0, err
}

//（查）获取所有的Namespace
func ListNamespace() ([]*Namespace, error) {
	var list []*Namespace
	o := orm.NewOrm()
	//_, err := o.QueryTable(&Client{}).OrderBy("-id").All(&list)
	_, err := o.QueryTable(&Namespace{}).OrderBy("-id").All(&list)
	return list, err
}
