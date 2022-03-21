package web2

import (
	"GoMessage/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

// 消息模板接口
type TemplateControllers struct {
	beego.Controller
}

// @Title /v1/web/template
// @Description 获取所有template
// @Success 200 响应成功
// @Failure 404 错误请求
// @router / [get]
func (this *TemplateControllers) GetAll() {
	listAllTemplate, err := models.ListAllTemplate()
	if err != nil {
		return
	}
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = listAllTemplate
	this.ServeJSON()
}

// @Title /v1/web/template
// @Description 获取指定编号的模板信息
// @Success 200 响应成功
// @Failure 404 错误请求
// @router /:namespace [get]
func (this *TemplateControllers) Get() {
	//获取get请求中的参数
	namespace := this.Ctx.Input.Param(":namespace")
	ns := models.GetNamespaceParamName(namespace)
	if ns != nil {
		temp := models.GetOneTemplateNid(ns.Id)
		//返回值
		this.Ctx.ResponseWriter.WriteHeader(200)
		this.Data["json"] = temp
		this.ServeJSON()
	} else {
		//返回值
		this.Ctx.ResponseWriter.WriteHeader(400)
		this.Data["json"] = "Namespace不存在"
		this.ServeJSON()
		return
	}

}

// @Title /v1/web/template
// @Description 添加用户模板
// @Success 200 响应成功
// @Failure 404 错误请求
// @router / [post]
func (this *TemplateControllers) Post() {
	type Param struct {
		RequestData struct {
			NamespaceName string            `json:"namespace_name"`
			Template      *models.Templates `json:"template"`
		} `json:"request_data"`
	}

	param := Param{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &param)

	//遍历所有的模板（按照namespace来遍历，只返回指定namespace下的遍历结果）
	for _, ttt := range models.ListNsTemplate(models.GetNamespaceParamName(param.RequestData.NamespaceName)) {
		//删除遍历到的结果
		models.DeleteTemplate(ttt)
	}

	//新增一个Template
	temp := models.ReadOrCreateTemplate(
		param.RequestData.Template.TemplateName,
		param.RequestData.Template.MessageTemplate,
		param.RequestData.Template.MessageMerge,
		models.GetNamespaceParamName(param.RequestData.NamespaceName),
	)

	//models.ReadOrCreateTemplate("bbb", "bbb", true, models.GetNamespaceParamName("test2"))
	//models.ReadOrCreateTemplate("ccc", "ccc", true, models.GetNamespaceParamName("test2"))
	//models.ReadOrCreateTemplate("ddd", "ddd", true, models.GetNamespaceParamName("test3"))
	//models.ReadOrCreateTemplate("eee", "eee", true, models.GetNamespaceParamName("test3"))
	//models.ReadOrCreateTemplate("fff", "fff", true, models.GetNamespaceParamName("test3"))
	//models.ReadOrCreateTemplate("fff", "fff", true, models.GetNamespaceParamName("test5"))

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = temp
	this.ServeJSON()
}
