package clients

import (
	"gomessage/apps/controllers/clients"
	"gomessage/apps/controllers/send"
	"gomessage/apps/models"
)

type ClientActionFeishu struct {
	Client *models.Client
}

func (c *ClientActionFeishu) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	if isMerge {
		msg := send.MessageJoint(contentList, "feishu")
		data := clients.PackFeishuMessage(client, msg)
		msgList = append(msgList, data)
	} else {
		for _, msg := range contentList {
			data := clients.PackFeishuMessage(client, msg)
			msgList = append(msgList, data)
		}
	}
	return msgList
}

func (c *ClientActionFeishu) PushMessages(messages []any) {
	url := send.RobotRandomUrl(c.Client.ExtendDingtalk.RobotUrlInfoList)
	for _, msg := range messages {
		send.Push(msg, url)
	}
}
