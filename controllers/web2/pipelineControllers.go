package web2

import (
	"GoMessage/controllers/piplineUtil"
	"GoMessage/models"
	"encoding/json"
	"fmt"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

//控制器：多管道入口
type PipelineControllers struct {
	beego.Controller
}

func (this *PipelineControllers) Post() {
	//获取namespace标识
	var requestNamespace string
	if this.Ctx.Request.RequestURI == "/gomessage" {
		//fmt.Println("是default入口")
		requestNamespace = "default"
	} else {
		//fmt.Println("是其它命名空间入口")
		requestNamespace = this.Ctx.Input.Param(":namespace")
	}
	fmt.Println(requestNamespace)

	//判断当前namespace是否是存在的（不为空时，才处理接下来的逻辑）
	oneNamespace := models.GetNamespaceParamName(requestNamespace)
	if oneNamespace != nil {
		//数据绑定
		var oneJson ArbitrarilyJsonData
		oneJson.RequestBody = this.Ctx.Input.RequestBody
		json.Unmarshal(oneJson.RequestBody, &oneJson.JsonData)
		oneJson.UpdateTime = time.Now()
		//fmt.Println(oneJson)
		//tmpDict := make(map[string]ArbitrarilyJsonData)
		//tmpDict[requestNamespace] = oneJson
		//ArrayCacheData.CacheDict = tmpDict
		ArrayCacheData.CacheDict[requestNamespace] = oneJson

		//拿到推送过来的数据
		//拿到推送过来的[]byte消息（OK）
		messageByteData := oneJson.RequestBody

		//从数据库中拿到"模板变量"映射关系 （数据库方法）
		//从数据库中拿到"模板内容" （数据库方法）
		//从数据库中拿到钉钉客户端的配置 （数据库方法）
		userConfig := piplineUtil.GetUserConfig(oneNamespace)

		//获得自定义变量与数据的映射 AnalysisData()
		analysisDataList := piplineUtil.AnalysisData(userConfig.ConfigMap, messageByteData)

		//得到渲染完成后的"消息体"的列表，还没有区分客户端；《消息体 + 客户端上下文配置》才等于最终需要推送出去的"完整消息"
		templateMessageList := piplineUtil.TemplateRenders(userConfig.MessageTemplate, analysisDataList)

		//判断消息是否聚合发送
		if userConfig.MessageMerge {
			MergeSend(templateMessageList, userConfig)
		} else {
			DisperseSend(templateMessageList, userConfig)
		}

		//单通道的流程处理完成后，返回200状态
		this.Ctx.ResponseWriter.WriteHeader(200)
		this.Data["json"] = "ok"
		this.ServeJSON()

	} else {
		//如果Namespace不存在，则返回400错误
		this.Ctx.ResponseWriter.WriteHeader(400)
		this.Data["json"] = "Namespace不存在"
		this.ServeJSON()
	}

} // post end

//========================
//消息（聚合）发送，所有消息拼接成一个整体
//========================
func MergeSend(messageList []string, userConfig piplineUtil.Config) {
	for _, ccc := range userConfig.ActiveClient {
		oneClientInfo, _ := models.GetClient(ccc.Id)
		cType := oneClientInfo.Type

		if cType == "feishu" {
			msg := piplineUtil.MessageJoint(messageList, cType) //完成消息拼接，不同的客户端消息拼接格式不同
			url := piplineUtil.RobotRandomUrl(oneClientInfo.ClientFeishu.RobotUrlList)
			data := piplineUtil.MessageRendersFeishu(oneClientInfo, msg)
			piplineUtil.SendMessage(data, url)

		} else if cType == "dingtalk" {
			msg := piplineUtil.MessageJoint(messageList, cType)
			url := piplineUtil.RobotRandomUrl(oneClientInfo.ClientDingtalk.RobotUrlList)
			data := piplineUtil.MessageRendersDingtalk(oneClientInfo.ClientDingtalk.RobotKeyword, msg)
			piplineUtil.SendMessage(data, url)

		} else if cType == "wechat" {
			msg := piplineUtil.MessageJoint(messageList, cType)
			wechat := piplineUtil.WeChat{
				Corpid:      oneClientInfo.ClientWechat.Corpid,
				Agentid:     oneClientInfo.ClientWechat.Agentid,
				Agentsecret: oneClientInfo.ClientWechat.Secret,
				Touser:      oneClientInfo.ClientWechat.Touser,
				Content:     msg,
			}
			wechat.PushMessage()
		}

		//data, url := WebMessageSend.VerdictClient(oneClientInfo, msg)
		//fmt.Println(data, url)
		//WebMessageSend.SendMessage(data, url)
	}

}

//========================
//消息（分散）发送，切割成一条一条的发送出去
//========================
func DisperseSend(messageList []string, userConfig piplineUtil.Config) {
	for _, oneMessage := range messageList {
		for _, oneClient := range userConfig.ActiveClient {
			oneClientInfo, _ := models.GetClient(oneClient.Id)
			cType := oneClientInfo.Type

			if cType == "feishu" {
				url := piplineUtil.RobotRandomUrl(oneClientInfo.ClientFeishu.RobotUrlList)
				data := piplineUtil.MessageRendersFeishu(oneClientInfo, oneMessage)
				piplineUtil.SendMessage(data, url)

			} else if cType == "dingtalk" {
				url := piplineUtil.RobotRandomUrl(oneClientInfo.ClientDingtalk.RobotUrlList)
				data := piplineUtil.MessageRendersDingtalk(oneClientInfo.ClientDingtalk.RobotKeyword, oneMessage)
				piplineUtil.SendMessage(data, url)

			} else if cType == "wechat" {
				fmt.Println("进入wechat判断...")
				wechat := piplineUtil.WeChat{
					Corpid:      oneClientInfo.ClientWechat.Corpid,
					Agentid:     oneClientInfo.ClientWechat.Agentid,
					Agentsecret: oneClientInfo.ClientWechat.Secret,
					Touser:      oneClientInfo.ClientWechat.Touser,
					Content:     oneMessage,
				}
				wechat.PushMessage()
			}
		}
	}
	fmt.Println("非聚合的方式发送完成...（已经完成for循环，或没有进入for循环）")
}
