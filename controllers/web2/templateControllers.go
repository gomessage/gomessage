package web2

import (
	"GoMessage/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

//控制器：用户消息体的模板
type TemplateControllers struct {
	beego.Controller
}

// @Title Get user template
// @Description Get user template
// @Success 200 响应成功
// @Failure 404 错误请求
// @router /template [get]
func (this *TemplateControllers) Get() {
	t := models.GetOneTemplate("default")
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = t
	this.ServeJSON()
}

// @Title /v1/web/template
// @Description 添加用户模板
// @Success 200 响应成功
// @Failure 404 错误请求
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

	temp := models.ReadOrCreateTemplate("default", param.RequestData.MessageTemplate, param.RequestData.MessageMerge, models.GetNamespace("default"))
	models.ReadOrCreateTemplate("bbb", "bbb", true, models.GetNamespace("test2"))
	models.ReadOrCreateTemplate("ccc", "ccc", true, models.GetNamespace("test2"))
	models.ReadOrCreateTemplate("ddd", "ddd", true, models.GetNamespace("test3"))
	models.ReadOrCreateTemplate("eee", "eee", true, models.GetNamespace("test3"))
	models.ReadOrCreateTemplate("fff", "fff", true, models.GetNamespace("test3"))
	models.ReadOrCreateTemplate("fff", "fff", true, models.GetNamespace("test55"))

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = temp
	this.ServeJSON()
}
