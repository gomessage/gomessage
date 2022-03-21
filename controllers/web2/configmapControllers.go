package web2

import (
	"GoMessage/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

//映射关系接口
type MapControllers struct {
	beego.Controller
}

// @Title /v1/web/map
// @Description 获取所有用户变量
// @Success 200 响应成功
// @Failure 404 错误请求
// @router / [get]
func (this *MapControllers) GetAll() {
	listAllConfigMap, err := models.ListAllConfigMap()
	if err != nil {
		return
	}
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = listAllConfigMap
	this.ServeJSON()
}

// @Title /v1/web/map
// @Description 获取指定namespace下的用户变量
// @Success 200 响应成功
// @Failure 404 错误请求
// @router /:namespace [get]
func (this *MapControllers) Get() {
	namespace := this.Ctx.Input.Param(":namespace")
	mList := models.ListNsConfigMap(models.GetNamespaceParamName(namespace))
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = mList
	this.ServeJSON()
}

// @Title /v1/web/map
// @Description 创建一组新的用户变量（同一Namespace下的老变量会被删掉，一个Namespace下永远只有一组用户变量）
// @Param key_value_list body string false "存放用户变量的一个list"
// @Success 200 {object} []models.Configmaps
// @Failure 404 错误请求
// @router / [post]
func (this *MapControllers) Post() {
	//解析request中的数据结构
	type Param struct {
		NamespaceName string              `json:"namespace_name"`
		KeyValueList  []map[string]string `json:"key_value_list"`
	}
	//获取request中的数据
	param := Param{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &param)

	//创建用户变量时，判断namespace是否已经存在
	namespace := models.GetNamespaceParamName(param.NamespaceName)
	if namespace != nil {

		//先删除当前namespace下的所有用户变量
		for _, cc := range models.ListNsConfigMap(namespace) {
			models.DeleteConfigMap(&cc, namespace)
		}

		//新建一个切片，暂存需要返回给用户的数据
		var ResponseData []models.Configmaps
		//批量写入新的配置
		for _, oneKeyValue := range param.KeyValueList {
			key := oneKeyValue["key"]
			value := oneKeyValue["value"]
			result := models.ReadOrCreateMap(key, value, namespace)
			ResponseData = append(ResponseData, result)
		}
		//返回值
		this.Ctx.ResponseWriter.WriteHeader(200)
		this.Data["json"] = ResponseData
		this.ServeJSON()
	} else {
		//返回值
		this.Ctx.ResponseWriter.WriteHeader(400)
		this.Data["json"] = "Namespace不存在"
		this.ServeJSON()
		return
	}

}
