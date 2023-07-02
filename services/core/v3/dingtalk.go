package v3

import (
	"gomessage/models"
	"gomessage/pkg/utils"
	"gomessage/services/core/v1"
	"gomessage/services/format"
)

type ClientActionDingtalk struct {
	Client *models.Client
}

func (c *ClientActionDingtalk) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	if isMerge {
		msg := v1.MessageJoint(contentList, utils.VarDingtalk)
		data := format.PackDingtalkMessage(client.ExtendDingtalk.RobotKeyword, msg)
		msgList = append(msgList, data)
	} else {
		for _, msg := range contentList {
			data := format.PackDingtalkMessage(client.ExtendDingtalk.RobotKeyword, msg)
			msgList = append(msgList, data)
		}
	}
	return msgList
}

func (c *ClientActionDingtalk) PushMessages(messages []any) {
	url := v1.RobotRandomUrl(c.Client.ExtendDingtalk.RobotUrlRandomList)
	for _, msg := range messages {
		v1.Push(msg, url)
	}
}
