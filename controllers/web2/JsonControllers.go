package alertmanager

import (
	"GoMessage/client/dingtalk"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

//全局变量，其它文件中可以写入，而我这个接口可以读取
var TmpJsonData ResponseJsonData

type ResponseJsonData struct {
	JsonData   dingtalk.Messages `json:"json_data"`
	UpdateTime time.Time         `json:"update_time"`
}

type JsonControllers struct {
	beego.Controller
}

// @Title 推送数据
// @Description 推送数据到钉钉，数据来自于Alertmanager的webhook推送
// @Success 200 ok
// @router /json [get]
func (this *JsonControllers) Get() {
	//tmpJsonData := ResponseJsonData{}

	fmt.Println(TmpJsonData)

	//bs, _ := json.Marshal(TmpJsonData)
	//var out bytes.Buffer
	//json.Indent(&out, bs, "", "\t")
	//fmt.Printf("aaa=%v\n", out.String())

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = TmpJsonData
	this.ServeJSON()
}
