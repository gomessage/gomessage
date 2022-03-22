package bak

//
//import (
//	"GoMessage/controllers/piplineUtil"
//	"encoding/json"
//	beego "github.com/beego/beego/v2/server/web"
//	"time"
//)
//
////控制器：信息实际进入接口
//type ApiControllers struct {
//	beego.Controller
//}
//
//func (this *ApiControllers) Get() {
//	//给出返回值
//	this.Ctx.ResponseWriter.WriteHeader(200)
//	this.Data["json"] = "GoMessage"
//	this.ServeJSON()
//}
//
//func (this *ApiControllers) Post() {
//	//进行CacheData数据绑定
//	CacheData.RequestBody = this.Ctx.Input.RequestBody
//	json.Unmarshal(CacheData.RequestBody, &CacheData.JsonData)
//	CacheData.UpdateTime = time.Now()
//
//	//拿到推送过来的[]byte消息（OK）
//	messageByteData := CacheData.RequestBody
//
//	//从数据库中拿到"模板变量"映射关系 （数据库方法）
//	//从数据库中拿到"模板内容" （数据库方法）
//	//从数据库中拿到钉钉客户端的配置 （数据库方法）
//	userConfig := piplineUtil.GetUserConfig()
//
//	//获得自定义变量与数据的映射 AnalysisData()
//	analysisDataList := piplineUtil.AnalysisData(userConfig.ConfigMap, messageByteData)
//	//fmt.Println(analysisDataList)
//
//	//得到渲染完成后的模板消息列表，测试还没有根据不同的客户端渲染成完整的推送消息体
//	templateMessageList := piplineUtil.TemplateRenders(userConfig.MessageTemplate, analysisDataList)
//	//fmt.Println("-------------", len(templateMessageList))
//
//	//判断消息是否聚合发送
//	if userConfig.MessageMerge {
//		MergeSend(templateMessageList, userConfig)
//	} else {
//		DisperseSend(templateMessageList, userConfig)
//	}
//
//	//给出返回值
//	this.Ctx.ResponseWriter.WriteHeader(200)
//	this.Data["json"] = "ok"
//	this.ServeJSON()
//}
