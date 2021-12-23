package web2

import (
	"GoMessage/apps/WebMessageSend"
	"GoMessage/models"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type ApiControllers struct {
	beego.Controller
}

func (this *ApiControllers) Get() {
	//给出返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = "GoMessage"
	this.ServeJSON()
}

func (this *ApiControllers) Post() {
	//进行CacheData数据绑定
	CacheData.RequestBody = this.Ctx.Input.RequestBody
	json.Unmarshal(CacheData.RequestBody, &CacheData.MessageData)
	CacheData.UpdateTime = time.Now()

	//拿到推送过来的[]byte消息（OK）
	messageByteData := CacheData.RequestBody

	//从数据库中拿到"模板变量"映射关系 （数据库方法）
	//从数据库中拿到"模板内容" （数据库方法）
	//从数据库中拿到钉钉客户端的配置 （数据库方法）
	userConfig := WebMessageSend.GetUserConfig()

	//获得自定义变量与数据的映射 AnalysisData()
	analysisDataList := WebMessageSend.AnalysisData(userConfig.ConfigMap, messageByteData)
	//fmt.Println(analysisDataList)

	//得到渲染完成后的模板消息列表，测试还没有根据不同的客户端渲染成完整的推送消息体
	templateMessageList := WebMessageSend.TemplateRenders(userConfig.MessageTemplate, analysisDataList)
	//fmt.Println("-------------", len(templateMessageList))

	//判断消息是否聚合发送
	if userConfig.MessageMerge {
		MergeSend(templateMessageList, userConfig)
	} else {
		DisperseSend(templateMessageList, userConfig)
	}

	//给出返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = "ok"
	this.ServeJSON()
}

//========================
//消息（聚合）发送，所有消息拼接成一个整体
//========================
func MergeSend(messageList []string, userConfig WebMessageSend.Config) {
	for _, ccc := range userConfig.ActiveClient {
		oneClientInfo, _ := models.GetClient(ccc.Id)
		cType := oneClientInfo.Type

		if cType == "feishu" {
			msg := WebMessageSend.MessageJoint(messageList, cType) //完成消息拼接，不同的客户端消息拼接格式不同
			url := WebMessageSend.RobotRandomUrl(oneClientInfo.ClientFeishu.RobotUrlList)
			data := WebMessageSend.MessageRendersFeishu(oneClientInfo, msg)
			WebMessageSend.SendMessage(data, url)

		} else if cType == "dingtalk" {
			msg := WebMessageSend.MessageJoint(messageList, cType)
			url := WebMessageSend.RobotRandomUrl(oneClientInfo.ClientDingtalk.RobotUrlList)
			data := WebMessageSend.MessageRendersDingtalk(oneClientInfo.ClientDingtalk.RobotKeyword, msg)
			WebMessageSend.SendMessage(data, url)

		} else if cType == "wechat" {
			msg := WebMessageSend.MessageJoint(messageList, cType)
			wechat := WebMessageSend.WeChat{
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
func DisperseSend(messageList []string, userConfig WebMessageSend.Config) {
	for _, oneMessage := range messageList {
		for _, oneClient := range userConfig.ActiveClient {
			oneClientInfo, _ := models.GetClient(oneClient.Id)
			cType := oneClientInfo.Type

			if cType == "feishu" {
				url := WebMessageSend.RobotRandomUrl(oneClientInfo.ClientFeishu.RobotUrlList)
				data := WebMessageSend.MessageRendersFeishu(oneClientInfo, oneMessage)
				WebMessageSend.SendMessage(data, url)

			} else if cType == "dingtalk" {
				url := WebMessageSend.RobotRandomUrl(oneClientInfo.ClientDingtalk.RobotUrlList)
				data := WebMessageSend.MessageRendersDingtalk(oneClientInfo.ClientDingtalk.RobotKeyword, oneMessage)
				WebMessageSend.SendMessage(data, url)

			} else if cType == "wechat" {
				fmt.Println("进入wechat判断...")
				wechat := WebMessageSend.WeChat{
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
	fmt.Println("没有进入for循环")
}
