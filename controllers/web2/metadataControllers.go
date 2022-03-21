package web2

import (
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

// 元数据接口
type JsonControllers struct {
	beego.Controller
}

// @Title /v1/web/json
// @Description 劫持到过境数据
// @Success 200 响应成功
// @Failure 404 错误请求
// @router /json/ [get]
func (this *JsonControllers) Get() {
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = CacheData
	this.ServeJSON()
}

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
