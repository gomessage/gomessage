package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(Json))
}

type Json struct {
	Id         int       `orm:"pk;auto"`                     //id
	Key        string    `orm:"size(500)"`                   //配置的key
	Value      string    `orm:"size(500)"`                   //配置的value
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"` //创建时间
}

//######################
//函数：添加用户
//######################
func AddMap(key string, value string) int {
	//密码加密

	o := orm.NewOrm()
	j := new(Json)
	j.Key = key
	j.Value = value

	uid, err := o.Insert(j)
	if err != nil {
		fmt.Errorf("数据插入失败")
	}
	return int(uid)
}

//######################
//函数：查询和创建一个用户，确保它始终存在
//######################
func ReadOrCreateMap(key string, value string) Json {
	o := orm.NewOrm()
	json := Json{}
	json.Key = key
	json.Value = value
	newCreate, id, err := o.ReadOrCreate(&json, "Id")
	if err != nil {
		panic(err)
	}
	//判断用户是否是新创建的
	if newCreate {
		fmt.Println("新创建用户ID：", id)
	} else {
		fmt.Println("查询到用户ID：", id)
		//如果用户存在，则更新用户的信息
		_, err := o.Update(&json)
		if err != nil {
			panic(err)
		}
	}
	return json
}

//######################
//函数：查询所有用户
//######################
func QueryAllMap() []Json {
	var jsons []Json
	json := Json{}

	o := orm.NewOrm()
	num, _ := o.QueryTable(&json).All(&jsons)
	fmt.Printf("受影响的行数：%v\n", num)

	return jsons
}

//######################
//函数：删除用户
//######################
func DeleteMap(j Json) int64 {
	o := orm.NewOrm()
	//json := Json{Id: Id}
	num, err2 := o.Delete(&j)
	if err2 != nil {
		fmt.Println(err2)
	}
	return num
}
