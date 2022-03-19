package web2

import (
	"GoMessage/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

//控制器：命名空间管理
type NamespaceControllers struct {
	beego.Controller
}

// @Title /v1/web/namespace
// @Description 获取所有命名空间
// @Param request_data body models.Namespaces true "Namespace对象"
// @Success 200 响应成功
// @Failure 404 错误请求
// @router /namespace [get]
func (this *NamespaceControllers) GetAll() {
	listNamespace, err := models.ListNamespace()
	if err != nil {
		return
	}
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = listNamespace
	this.ServeJSON()
}

// @Title /v1/web/namespace
// @Description 添加一个命名空间
// @Param request_data body models.Namespaces true "Namespace对象"
// @Success 200 响应成功
// @Failure 404 错误请求
// @router /namespace [post]
func (this *NamespaceControllers) Post() {
	//request中的数据结构
	type Param struct {
		RequestData struct {
			*models.Namespaces
		} `json:"request_data"`
	}

	//获取request中的数据
	param := Param{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &param)
	if err != nil {
		return
	}

	newNamespace := models.Namespaces{
		Name:        param.RequestData.Name,
		Description: param.RequestData.Description,
	}

	namespaceId, err := models.AddNamespace(&newNamespace)
	if err != nil {
		return
	}

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = struct {
		Message string `json:"message"`
		Id      int64  `json:"id"`
	}{"Namespace添加成功", namespaceId}
	this.ServeJSON()
}

// @Title /v1/web/namespace
// @Description 删除一个命名空间
// @Param request_data body models.Namespaces true "Namespace对象"
// @Success 200 响应成功
// @Failure 404 错误请求
// @router /namespace [delete]
func (this *NamespaceControllers) Delete() {
	//request中的数据结构
	type Param struct {
		RequestData struct {
			*models.Namespaces
		} `json:"request_data"`
	}

	//获取request中的数据
	param := Param{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &param)
	if err != nil {
		return
	}

	//获取namespace对象
	ns := models.GetNamespaceParamName(param.RequestData.Name)

	//删除templates表中关联数据
	for _, temp := range models.ListNsTemplate(ns) {
		models.DeleteTemplate(temp)
	}

	//删除configmaps表中关联的数据
	for _, configmap := range models.ListNsConfigMap(ns) {
		models.DeleteOneConfigMap(&configmap)
	}

	//最后再删除命名空间
	namespaceId, delErr := models.DelNamespace(param.RequestData.Name)
	if delErr != nil {
		return
	}

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = struct {
		Message string `json:"message"`
		Id      int64  `json:"id"`
	}{"Namespace删除成功", namespaceId}
	this.ServeJSON()
}

// @Title /v1/web/namespace
// @Description 修改一个命名空间
// @Param request_data body models.Namespaces true "Namespace对象"
// @Success 200 响应成功
// @Failure 404 错误请求
// @router /namespace [put]
func (this *NamespaceControllers) Put() {
	//request中的数据结构
	type Param struct {
		RequestData struct {
			*models.Namespaces
		} `json:"request_data"`
	}

	//获取request中的数据
	param := Param{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &param)
	if err != nil {
		return
	}

	oneNamespace := models.Namespaces{
		Id:          param.RequestData.Id,
		Name:        param.RequestData.Name,
		Description: param.RequestData.Description,
	}

	updateNum, err := models.UpdateNamespace(&oneNamespace)
	if err != nil {
		return
	}

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = struct {
		Message string `json:"message"`
		Num     int64  `json:"num"`
	}{
		"受影响的行数",
		updateNum,
	}
	this.ServeJSON()
}
