package v3

import (
	"gomessage/src/models"
	"gomessage/src/pkg/utils"
	v12 "gomessage/src/services/core/v1"
	"gomessage/src/services/format"
)

type ClientActionDingtalk struct {
	Client *models.Client
}

func (c *ClientActionDingtalk) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	if isMerge {
		msg := v12.MessageJoint(contentList, utils.VarDingtalk)
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
	url := v12.RobotRandomUrl(c.Client.ExtendDingtalk.RobotUrlRandomList)
	for _, msg := range messages {
		v12.Push(msg, url)
	}
}
