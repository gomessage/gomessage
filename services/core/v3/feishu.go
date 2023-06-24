package v3

import (
	"gomessage/models"
	"gomessage/services/clientFormat"
	v12 "gomessage/services/core/v1"
)

type ClientActionFeishu struct {
	Client *models.Client
}

func (c *ClientActionFeishu) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	if isMerge {
		msg := v12.MessageJoint(contentList, "feishu")
		data := clientFormat.PackFeishuMessage(client, msg)
		msgList = append(msgList, data)
	} else {
		for _, msg := range contentList {
			data := clientFormat.PackFeishuMessage(client, msg)
			msgList = append(msgList, data)
		}
	}
	return msgList
}

func (c *ClientActionFeishu) PushMessages(messages []any) {
	url := v12.RobotRandomUrl(c.Client.ExtendFeishu.RobotUrlInfoList)
	for _, msg := range messages {
		v12.Push(msg, url)
	}
}
