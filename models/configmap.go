package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(Configmap))
}

type Configmap struct {
	Id         int         `orm:"pk;auto"`                     //id
	Namespace  *Namespaces `orm:"rel(fk);cascade"`             //外键连接Namespace表
	Key        string      `orm:"size(500)"`                   //配置的key
	Value      string      `orm:"size(500)"`                   //配置的value
	CreateTime time.Time   `orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateTime time.Time   `orm:"auto_now;type(datetime)"`     //修改时间
}

//######################
//函数：添加一条数据（按namespace创建）
//######################
func ReadOrCreateMap(key string, value string, nm *Namespaces) Configmap {
	o := orm.NewOrm()
	newConfigmap := Configmap{
		Key:       key,
		Value:     value,
		Namespace: nm,
	}
	//根据命名空间和变量名进行查询，如果两者都存在，则不进行创建
	newCreate, id, err := o.ReadOrCreate(&newConfigmap, "Key", "Namespace")
	if err != nil {
		panic(err)
	}
	//判断用户是否是新创建的
	if newCreate {
		fmt.Println("新建ConfigMap，ID：", id)
	} else {
		fmt.Println("已存在ConfigMap，ID：", id)
		//如果用户存在，则更新用户的信息
		_, err := o.Update(&newConfigmap, "Namespace")
		if err != nil {
			panic(err)
		}
	}
	return newConfigmap
}

//######################
//函数：查询所有（按namespace查询）
//######################
func ListConfigMap(namespace *Namespaces) []Configmap {
	var configMapList []Configmap
	o := orm.NewOrm()
	num, err := o.QueryTable("configmap").Filter("namespace__id", &namespace.Id).All(&configMapList)
	if err != nil {
		return nil
	}
	fmt.Printf("查询全部ConfigMap，受影响的行数：%v\n", num)
	return configMapList
}

//######################
//函数：删除（按namespace删除）
//######################
func DeleteConfigMap(configmap *Configmap, namespace *Namespaces) int64 {
	o := orm.NewOrm()
	num, err := o.QueryTable("configmap").Filter("key", &configmap.Key).Filter("namespace__id", &namespace.Id).Delete()
	if err != nil {
		return 0
	}
	return num
}
