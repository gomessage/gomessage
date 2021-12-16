package web2

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

//========================
//可接受任意json格式的结构体
//========================
type ArbitrarilyJsonData struct {
	MessageData map[string]interface{} `json:"json_data"`
	UpdateTime  time.Time              `json:"update_time"`
	RequestBody []byte
}

//========================
//全局变量，其它文件中可以写入和读取
//========================
var CacheData ArbitrarilyJsonData

//========================
// JsonControllers控制器
//========================
type JsonControllers struct {
	beego.Controller
}

// @Title 推送数据
// @Description 推送数据到钉钉，数据来自于Alertmanager的webhook推送
// @Success 200 ok
// @router /json [get]
func (this *JsonControllers) Get() {
	//tmpJsonData := ArbitrarilyJsonData{}

	fmt.Println(CacheData)

	//bs, _ := json.Marshal(CacheData)
	//var out bytes.Buffer
	//json.Indent(&out, bs, "", "\t")
	//fmt.Printf("aaa=%v\n", out.String())

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = CacheData
	this.ServeJSON()
}
