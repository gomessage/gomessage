package v3

import (
	"gomessage/models"
	"gomessage/pkg/utils"
	"gomessage/services/core/v1"
	"gomessage/services/format"
)

type ClientActionFeishu struct {
	Client *models.Client
}

func (c *ClientActionFeishu) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	if isMerge {
		msg := v1.MessageJoint(contentList, utils.VarFeishu)
		data := format.PackFeishuMessage(client, msg)
		msgList = append(msgList, data)
	} else {
		for _, msg := range contentList {
			data := format.PackFeishuMessage(client, msg)
			msgList = append(msgList, data)
		}
	}
	return msgList
}

func (c *ClientActionFeishu) PushMessages(messages []any) {
	url := v1.RobotRandomUrl(c.Client.ExtendFeishu.RobotUrlRandomList)
	for _, msg := range messages {
		v1.Push(msg, url)
	}
}
