package alertmanager

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type JsonControllers struct {
	beego.Controller
}

// @Title 推送数据
// @Description 推送数据到钉钉，数据来自于Alertmanager的webhook推送
// @Success 200 ok
// @router /json [get]
func (this *JsonControllers) Get() {
	//tmpJsonData := ResponseJsonData{}

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

//只能接收固定json格式
//type ResponseJsonData struct {
//	JsonData   dingtalk.Messages `json:"json_data"`
//	UpdateTime time.Time         `json:"update_time"`
//}

//可以接受任意json格式
type ResponseJsonData struct {
	JsonData   map[string]interface{} `json:"json_data"`
	UpdateTime time.Time              `json:"update_time"`
}

//全局变量，其它文件中可以写入，而我这个接口可以读取
var CacheData ResponseJsonData
