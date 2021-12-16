package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type TestController struct {
	beego.Controller
}

// @router / [get]
func (this *TestController) Get() {

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = "hello"
	this.ServeJSON()
}

// @router / [post]
func (this *TestController) Post() {

	var CacheData map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &CacheData)
	fmt.Println(CacheData)

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = CacheData
	this.ServeJSON()
}
