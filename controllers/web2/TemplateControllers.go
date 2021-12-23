package web2

import (
	"GoMessage/models"
	"encoding/json"
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
	type Param struct {
		RequestData struct {
			*models.Templates
		} `json:"request_data"`
	}

	param := Param{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &param)

	//遍历删除全部配置存储
	for _, ttt := range models.QueryAllTemplate() {
		models.DeleteTemplate(ttt)
	}

	temp := models.ReadOrCreateTemplate("default", param.RequestData.MessageTemplate, param.RequestData.MessageMerge)

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = temp
	this.ServeJSON()
}
