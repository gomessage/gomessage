package web2

import (
	"GoMessage/models"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type TemplateControllers struct {
	beego.Controller
}

// @router /template [get]
func (this *TemplateControllers) Get() {
	t := models.GetOneTemplate("default")
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = t
	this.ServeJSON()
}

// @router /template [post]
func (this *TemplateControllers) Post() {
	type requestData struct {
		MessageTemplate string `json:"message_template"`
	}
	data := requestData{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &data)

	//遍历删除全部配置存储
	for _, ttt := range models.QueryAllTemplate() {
		models.DeleteTemplate(ttt)
	}

	fmt.Println(data)
	temp := models.ReadOrCreateTemplate("default", data.MessageTemplate)

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = temp
	this.ServeJSON()
}
