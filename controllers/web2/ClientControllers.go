package web2

import (
	"GoMessage/models"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type ClientControllers struct {
	beego.Controller
}

// @router /client [get]
func (this *ClientControllers) Get() {
	mList := models.QueryAllDingtalkClient()
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = mList
	this.ServeJSON()
}

// @router /client [post]
func (this *ClientControllers) Post() {
	type Data struct {
		RequestData []map[string]string `json:"request_data"`
	}
	data := Data{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)

	//删除所有数据
	allData := models.QueryAllDingtalkClient()
	for _, st := range allData {
		o := orm.NewOrm()
		_, err := o.Delete(&st)
		if err != nil {
			return
		}
	}

	//批量添加数据（单条插入，多次循环）
	for _, dict := range data.RequestData {
		m := models.DingtalkClient{
			RobotName:    dict["robotName"],
			RobotUrl:     dict["robotUrl"],
			RobotKeyword: dict["robotKeyword"],
		}

		id := models.AddOneDingtalkClient(m)
		fmt.Println("数据ID：", id)
	}

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = "hello"
	this.ServeJSON()
}
