package alertmanager

import (
	"GoMessage/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

//全局变量，其它文件中可以写入，而我这个接口可以读取
//var CacheData ResponseJsonData

//type ResponseJsonData struct {
//	JsonData   dingtalk.Messages `json:"json_data"`
//	UpdateTime time.Time         `json:"update_time"`
//}

type MapControllers struct {
	beego.Controller
}

// Post
// @Title 推送数据
// @Description 推送数据到钉钉，数据来自于Alertmanager的webhook推送
// @Success 200 ok
// @router /map [post]
func (this *MapControllers) Post() {
	type requestData struct {
		KeyValueList []map[string]string `json:"key_value_list"`
	}
	r := requestData{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &r)
	if err != nil {
		panic(err)
	}

	//遍历删除全部配置存储
	for _, json := range models.GetAllUsers() {
		models.DeleteUser(json)
		//fmt.Println(result)
	}

	//新建一个切片，暂存需要返回给用户的数据
	var ResponseJsons []models.Json

	//批量写入新的配置
	for _, json := range r.KeyValueList {
		for key, vv := range json {
			result := models.ReadOrCreateUser(key, vv)
			ResponseJsons = append(ResponseJsons, result)
		}
	}

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = ResponseJsons
	this.ServeJSON()
}
