package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(Configmaps))
}

type Configmaps struct {
	Id          int       `json:"id" orm:"pk;auto"`                    //id
	NamespaceId int       `json:"namespace_id"`                        //外键连接Namespace表
	Key         string    `json:"key" orm:"size(500)"`                 //配置的key
	Value       string    `json:"value" orm:"size(500)"`               //配置的value
	CreateTime  time.Time `json:"-" orm:"auto_now_add;type(datetime)"` //创建时间
	UpdateTime  time.Time `json:"-" orm:"auto_now;type(datetime)"`     //修改时间
}

//######################
//函数：添加一条数据（按namespace创建）
//######################
func ReadOrCreateMap(key string, value string, nm *Namespaces) Configmaps {
	o := orm.NewOrm()
	newConfigmap := Configmaps{
		NamespaceId: nm.Id,
		Key:         key,
		Value:       value,
	}
	//根据命名空间和变量名进行查询，如果两者都存在，则不进行创建
	newCreate, id, err := o.ReadOrCreate(&newConfigmap, "key", "namespace_id")
	if err != nil {
		panic(err)
	}
	//判断用户是否是新创建的
	if newCreate {
		fmt.Println("新建ConfigMap，ID：", id)
	} else {
		fmt.Println("已存在ConfigMap，ID：", id)
		//如果用户存在，则更新用户的信息
		_, err := o.Update(&newConfigmap, "namespace_id")
		if err != nil {
			panic(err)
		}
	}
	return newConfigmap
}

//######################
//函数：删除（按namespace条件+configmap条件进行删除）
//######################
func DeleteConfigMap(configmap *Configmaps, ns *Namespaces) int64 {
	o := orm.NewOrm()
	num, err := o.QueryTable("configmaps").Filter("key", &configmap.Key).Filter("namespace_id", &ns.Id).Delete()
	if err != nil {
		return 0
	}
	return num
}

//######################
//函数：删除（按namespace条件+configmap条件进行删除）
//######################
func DeleteOneConfigMap(configmap *Configmaps) int64 {
	o := orm.NewOrm()
	num, err := o.QueryTable("configmaps").Filter("key", &configmap.Key).Delete()
	if err != nil {
		return 0
	}
	return num
}

//######################
//函数：查询所有（按namespace查询）
//######################
func ListNsConfigMap(ns *Namespaces) []Configmaps {
	var configMapList []Configmaps
	o := orm.NewOrm()
	num, err := o.QueryTable("configmaps").Filter("namespace_id", &ns.Id).All(&configMapList)
	if err != nil {
		return nil
	}
	fmt.Printf("查询全部ConfigMap，受影响的行数：%v\n", num)
	return configMapList
}

//######################
//函数：查询所有（不加条件）
//######################
func ListAllConfigMap() ([]*Configmaps, error) {
	var list []*Configmaps
	o := orm.NewOrm()
	_, err := o.QueryTable(Configmaps{}).OrderBy("-id").All(&list)
	if err != nil {
		return nil, err
	}
	return list, err
}
