package v3

import (
	"gomessage/models"
	"gomessage/services/clientFormat"
	v12 "gomessage/services/core/v1"
)

type ClientActionDingtalk struct {
	Client *models.Client
}

func (c *ClientActionDingtalk) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	if isMerge {
		msg := v12.MessageJoint(contentList, "dingtalk")
		data := clientFormat.PackDingtalkMessage(client.ExtendDingtalk.RobotKeyword, msg)
		msgList = append(msgList, data)
	} else {
		for _, msg := range contentList {
			data := clientFormat.PackDingtalkMessage(client.ExtendDingtalk.RobotKeyword, msg)
			msgList = append(msgList, data)
		}
	}
	return msgList
}

func (c *ClientActionDingtalk) PushMessages(messages []any) {
	url := v12.RobotRandomUrl(c.Client.ExtendDingtalk.RobotUrlInfoList)
	for _, msg := range messages {
		v12.Push(msg, url)
	}
}
