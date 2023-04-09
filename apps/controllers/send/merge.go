package send

import (
	"fmt"
	"gomessage/apps/controllers/clients"
	"gomessage/apps/models"
)

// Merge 消息（聚合）发送，所有消息拼接成一个整体
func Merge(messageList []string, userConfig UserConfig) {
	for _, ccc := range userConfig.ActiveClient {
		oneClientInfo, _ := models.GetClientById(int(ccc.ID))
		cType := oneClientInfo.ClientType

		if cType == "feishu" {
			msg := MessageJoint(messageList, cType) //完成消息拼接，不同的客户端消息拼接格式不同
			url := RobotRandomUrl(oneClientInfo.ExtendFeishu.RobotUrls)
			data := clients.MessageRendersFeishu(oneClientInfo, msg)
			SendMessage(data, url)

		} else if cType == "dingtalk" {
			msg := MessageJoint(messageList, cType)
			url := RobotRandomUrl(oneClientInfo.ExtendDingtalk.RobotUrlInfoList)
			data := clients.MessageRendersDingtalk(oneClientInfo.ExtendDingtalk.RobotKeyword, msg)
			SendMessage(data, url)

		} else if cType == "wechat" {
			msg := MessageJoint(messageList, cType)
			wechat := clients.WeChat{
				Corpid:      oneClientInfo.ExtendWechat.CorpId,
				Agentid:     oneClientInfo.ExtendWechat.AgentId,
				Agentsecret: oneClientInfo.ExtendWechat.Secret,
				Touser:      oneClientInfo.ExtendWechat.Touser,
				Content:     msg,
			}
			wechat.PushMessage()
		}
	}

}

// Disperse 消息（分散）发送，切割成一条一条的发送出去
func Disperse(messageList []string, userConfig UserConfig) {
	for _, oneMessage := range messageList {
		for _, oneClient := range userConfig.ActiveClient {
			oneClientInfo, _ := models.GetClientById(int(oneClient.ID))
			cType := oneClientInfo.ClientType

			if cType == "feishu" {
				url := RobotRandomUrl(oneClientInfo.ExtendFeishu.RobotUrls)
				data := clients.MessageRendersFeishu(oneClientInfo, oneMessage)
				SendMessage(data, url)

			} else if cType == "dingtalk" {
				url := RobotRandomUrl(oneClientInfo.ExtendDingtalk.RobotUrlInfoList)                          //随机获取一个机器人地址
				data := clients.MessageRendersDingtalk(oneClientInfo.ExtendDingtalk.RobotKeyword, oneMessage) //拼装一个完整的服务钉钉消息推送规范的内容体
				SendMessage(data, url)

			} else if cType == "wechat" {
				fmt.Println("进入wechat判断...")
				wechat := clients.WeChat{
					Corpid:      oneClientInfo.ExtendWechat.CorpId,
					Agentid:     oneClientInfo.ExtendWechat.AgentId,
					Agentsecret: oneClientInfo.ExtendWechat.Secret,
					Touser:      oneClientInfo.ExtendWechat.Touser,
					Content:     oneMessage,
				}
				wechat.PushMessage()
			}
		}
	}
	fmt.Println("不切割信息发送，for循环已结束")
}
