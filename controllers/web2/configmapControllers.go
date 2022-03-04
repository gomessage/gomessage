package web2

import (
	"GoMessage/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

//控制器：用户变量功能
type MapControllers struct {
	beego.Controller
}

// @Title /v1/web/map
// @Description 获取用户变量
// @Success 200 响应成功
// @Failure 404 错误请求
// @router /map [get]
func (this *MapControllers) Get() {
	mList := models.QueryAllMap()
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = mList
	this.ServeJSON()
}

// @Title /v1/web/map
// @Description 更新用户变量
// @Param key_value_list body string false "存放用户变量的一个list"
// @Success 200 {object} []models.Configmap
// @Failure 404 错误请求
// @router /map [post]
func (this *MapControllers) Post() {
	//解析request中的数据结构
	type requestData struct {
		KeyValueList []map[string]string `json:"key_value_list"`
	}
	//获取request中的数据
	r := requestData{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &r)
	if err != nil {
		panic(err)
	}

	//遍历删除全部配置存储
	for _, json := range models.QueryAllMap() {
		models.DeleteMap(json)
		//fmt.Println(result)
	}

	//新建一个切片，暂存需要返回给用户的数据
	var ResponseJsons []models.Configmap

	//批量写入新的配置
	for _, json := range r.KeyValueList {
		for key, vv := range json {
			result := models.ReadOrCreateMap(key, vv)
			ResponseJsons = append(ResponseJsons, result)
		}
	}

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = ResponseJsons
	this.ServeJSON()
}
