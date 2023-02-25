package views

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gomessage/apps/controllers/cachex"
	"gomessage/apps/controllers/clients"
	"gomessage/apps/controllers/sendMessage"
	"gomessage/apps/models"
	"gomessage/utils/runLog"
	"io"
	"net/http"
	"time"
)

// GoMessageByPost
// @Tags gomessage
// @Router /go/:namespace [POST]
func GoMessageByPost(g *gin.Context) {
	namespace := g.Param("namespace")
	if namespace == "message" || namespace == "default" {
		namespace = "default"
	}
	runLog.Log.Info("当前命名空间为：", namespace)

	//获取请求数据
	cachex.CacheData.Time = time.Now()
	cachex.CacheData.RequestByte, _ = io.ReadAll(g.Request.Body)                 //g.Request.Body中的数据只能读取一次，是因为"流"的指针被移位了
	g.Request.Body = io.NopCloser(bytes.NewBuffer(cachex.CacheData.RequestByte)) //向g.Request.Body回写数据

	if err := g.ShouldBindJSON(&cachex.CacheData.RequestJson); err != nil {
		return
	}

	//把推送过来的数据写入缓存
	cachex.SetCacheData(namespace, cachex.CacheData)

	//从数据库中拿到用户当前用户在图形界面上配置的参数
	//userConfig := sendMessage.GetUserConfig(namespace)

	//创建过境数据与用户变量之间的映射
	//analysisDataList := sendMessage.AnalysisData(userConfig.VariablesMap, cachex.CacheData.RequestByte)

	//得到渲染完成后的消息列表
	//templateMessageList := sendMessage.CompleteMessage(userConfig.MessageTemplate, analysisDataList)

	//判断消息是否聚合发送
	//if userConfig.MessageMerge {
	//    MergeSend(templateMessageList, userConfig)
	//} else {
	//    DisperseSend(templateMessageList, userConfig)
	//}

	g.JSON(http.StatusOK, "ok")
}

// GoMessageByGet
// @Tags gomessage
// @Router /go/message [GET]
func GoMessageByGet(g *gin.Context) {
	namespace := g.Param("namespace")
	if namespace == "message" || namespace == "default" {
		namespace = "default"
	}
	runLog.Log.Info("当前命名空间为：", namespace)

	result, err := models.GetNamespaceByName(namespace)
	if err != nil {
		g.JSON(http.StatusBadRequest, "Namespace不存在...")
	} else {
		g.JSON(http.StatusOK, result)
	}
}

// MergeSend 消息（聚合）发送，所有消息拼接成一个整体
func MergeSend(messageList []string, userConfig sendMessage.Config) {
	for _, ccc := range userConfig.ActiveClient {
		oneClientInfo, _ := models.GetClientById(int(ccc.ID))
		cType := oneClientInfo.ClientType

		if cType == "feishu" {
			msg := sendMessage.MessageJoint(messageList, cType) //完成消息拼接，不同的客户端消息拼接格式不同
			url := sendMessage.RobotRandomUrl(oneClientInfo.ClientInfoFeishu.RobotUrls)
			data := clients.MessageRendersFeishu(oneClientInfo, msg)
			sendMessage.SendMessage(data, url)

		} else if cType == "dingtalk" {
			msg := sendMessage.MessageJoint(messageList, cType)
			url := sendMessage.RobotRandomUrl(oneClientInfo.ClientInfoDingtalk.RobotUrls)
			data := clients.MessageRendersDingtalk(oneClientInfo.ClientInfoDingtalk.RobotKeyword, msg)
			sendMessage.SendMessage(data, url)

		} else if cType == "wechat" {
			msg := sendMessage.MessageJoint(messageList, cType)
			wechat := clients.WeChat{
				Corpid:      oneClientInfo.ClientInfoWechat.CorpId,
				Agentid:     oneClientInfo.ClientInfoWechat.AgentId,
				Agentsecret: oneClientInfo.ClientInfoWechat.Secret,
				Touser:      oneClientInfo.ClientInfoWechat.Touser,
				Content:     msg,
			}
			wechat.PushMessage()
		}
	}

}

// DisperseSend 消息（分散）发送，切割成一条一条的发送出去
func DisperseSend(messageList []string, userConfig sendMessage.Config) {
	for _, oneMessage := range messageList {
		for _, oneClient := range userConfig.ActiveClient {
			oneClientInfo, _ := models.GetClientById(int(oneClient.ID))
			cType := oneClientInfo.ClientType

			if cType == "feishu" {
				url := sendMessage.RobotRandomUrl(oneClientInfo.ClientInfoFeishu.RobotUrls)
				data := clients.MessageRendersFeishu(oneClientInfo, oneMessage)
				sendMessage.SendMessage(data, url)

			} else if cType == "dingtalk" {
				url := sendMessage.RobotRandomUrl(oneClientInfo.ClientInfoDingtalk.RobotUrls)
				data := clients.MessageRendersDingtalk(oneClientInfo.ClientInfoDingtalk.RobotKeyword, oneMessage)
				sendMessage.SendMessage(data, url)

			} else if cType == "wechat" {
				fmt.Println("进入wechat判断...")
				wechat := clients.WeChat{
					Corpid:      oneClientInfo.ClientInfoWechat.CorpId,
					Agentid:     oneClientInfo.ClientInfoWechat.AgentId,
					Agentsecret: oneClientInfo.ClientInfoWechat.Secret,
					Touser:      oneClientInfo.ClientInfoWechat.Touser,
					Content:     oneMessage,
				}
				wechat.PushMessage()
			}
		}
	}
	fmt.Println("没有进入for循环")
}