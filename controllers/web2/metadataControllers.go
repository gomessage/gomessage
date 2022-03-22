package web2

import (
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

func init() {
	ArrayCacheData.CacheDict = make(map[string]ArbitrarilyJsonData)
}

//劫持数据的接口
type JsonControllers struct {
	beego.Controller
}

// @Title /v1/metadata/json/:namespace
// @Description 劫持到过境数据
// @Success 200 响应成功
// @Failure 404 错误请求
// @router /json/:namespace [get]
func (this *JsonControllers) Get() {
	ns := this.Ctx.Input.Param(":namespace")
	jsonData := ArrayCacheData.CacheDict[ns]

	//fmt.Println(ArrayCacheData.CacheDict)
	//fmt.Println(jsonData)

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = jsonData
	this.ServeJSON()
}

//========================
//存放任意json数据结构的结构体
//========================
type ArbitrarilyJsonData struct {
	JsonData    map[string]interface{} `json:"json_data"`   //这个值是给前端js解析使用的，后端用不到这个值
	UpdateTime  time.Time              `json:"update_time"` //这个值也是给前端js解析使用的，后端用不到这个值
	RequestBody []byte                 `json:"-"`           //这个值是为了减少后端的request的读写次数，为了实现一起读写全文可用而声明的，不会传递到前端去使用
}

//========================
//缓存队列结构体
//========================
type CacheQueue struct {
	CacheDict map[string]ArbitrarilyJsonData
}

//========================
//全局变量，其它文件中可以写入和读取
//========================
var CacheData ArbitrarilyJsonData //即将作废
var ArrayCacheData CacheQueue     //实例化一个缓存队列，用来存放一堆缓存数据，需要在上文的init生命周期阶段进行一次初始化
