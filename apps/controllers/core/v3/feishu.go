package v3

import (
	"gomessage/apps/controllers/clientFormats"
	"gomessage/apps/controllers/core/v1"
	"gomessage/apps/models"
)

type ClientActionFeishu struct {
	Client *models.Client
}

func (c *ClientActionFeishu) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	if isMerge {
		msg := v1.MessageJoint(contentList, "feishu")
		data := clientFormats.PackFeishuMessage(client, msg)
		msgList = append(msgList, data)
	} else {
		for _, msg := range contentList {
			data := clientFormats.PackFeishuMessage(client, msg)
			msgList = append(msgList, data)
		}
	}
	return msgList
}

func (c *ClientActionFeishu) PushMessages(messages []any) {
	url := v1.RobotRandomUrl(c.Client.ExtendDingtalk.RobotUrlInfoList)
	for _, msg := range messages {
		v1.Push(msg, url)
	}
}
