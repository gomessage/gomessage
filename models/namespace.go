package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

//命名空间（表结构）
type Namespaces struct {
	Id          int       `json:"id" orm:"pk;auto"`                    //id
	Name        string    `json:"name" orm:"size(50);unique"`          //命名空间
	Description string    `json:"description" orm:"size(200);null"`    //描述
	CreateTime  time.Time `json:"-" orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateTime  time.Time `json:"-" orm:"auto_now;type(datetime)"`     //修改时间
}

//注册数据表
func init() {
	orm.RegisterModel(new(Namespaces))
}

//创建默认命名空间，由main入口那里触发调用
func CreateDefaultNamespace() {
	//实例化一个默认的namespace对象，然后在开局的main函数中创造出来
	defaultNamespace := Namespaces{
		Name:        "default",
		Description: "Default Namespace",
	}
	//创建一个Namespace当做系统的default命名空间
	_, err := AddNamespace(&defaultNamespace)
	if err != nil {
		return
	}
}

//增：添加一个namespace
func AddNamespace(ns *Namespaces) (int64, error) {
	o := orm.NewOrm()
	//是否产生创建动作、对象的id、报错
	created, id, err := o.ReadOrCreate(ns, "Name")
	if err != nil {
		return 0, err
	} else {
		if created {
			fmt.Println("新增命名空间，Id:", id)
			return id, err
		} else {
			fmt.Println("查询到命名空间，Id:", id)
			return id, err
		}
	}
}

//删：删除一个namespace，并级联删除该命名空间下的其它资源
func DelNamespace(nsName string) (int64, error) {
	//只要在其它表的外键中，添加该内容`orm:"rel(fk);cascade"`，就会在删除本namespace数据时，被orm自动级联删除对方的数据：
	o := orm.NewOrm()
	i, err := o.QueryTable("namespaces").Filter("name", nsName).Delete()
	if err != nil {
		return 0, err
	}
	return i, err
}

//改：（这个改方法可能写的不正确，后面有空了再修改）
func UpdateNamespace(oleNs *Namespaces, newNs *Namespaces) (int64, error) {
	//var oneNs *Namespaces
	//var num2 int64
	//var err2 error
	//
	//o := orm.NewOrm()
	//err := o.QueryTable(&Namespaces{}).Filter("id", &ns.Id).One(&oneNs)
	//if err != nil {
	//	return 0, err
	//} else {
	//	oneNamespace := Namespaces{Id: ns.Id}
	//	if num2, err2 = o.Update(&oneNamespace); err == nil {
	//		fmt.Println("修改Namespace受影响的行数：", num2)
	//		return num2, err2
	//	}
	//	return 0, err2
	//}

	var num int64
	var err error
	o := orm.NewOrm()
	temp := Namespaces{Id: oleNs.Id}
	if o.Read(&temp) == nil {
		temp.Name = newNs.Name
		temp.Description = newNs.Description
		if num, err = o.Update(&temp); err == nil {
			fmt.Println(num)
		}
	}
	return num, err
}

//查：获取所有的Namespace
func ListNamespace() ([]*Namespaces, error) {
	var list []*Namespaces
	o := orm.NewOrm()
	_, err := o.QueryTable(&Namespaces{}).OrderBy("-id").All(&list)
	return list, err
}

//查：输入命名空间名称，就可以获得该namespace的完整对象。
func GetNamespaceParamName(nsName string) *Namespaces {
	var oneNamespace Namespaces
	o := orm.NewOrm()
	err := o.QueryTable("namespaces").Filter("name", nsName).One(&oneNamespace)
	if err != nil {
		//fmt.Println("命名空间" + nsName + "不存在~，向调用方返回nil...")
		return nil
	}
	return &oneNamespace
}

//查：输入命名空间的ID，就可以获得该namespace的完整对象。
func GetNamespaceParamId(nsId int) *Namespaces {
	var oneNamespace Namespaces
	o := orm.NewOrm()
	err := o.QueryTable("namespaces").Filter("id", nsId).One(&oneNamespace)
	if err != nil {
		//fmt.Println("命名空间" + namespaceName + "不存在~，向调用方返回nil...")
		return nil
	}
	return &oneNamespace
}
