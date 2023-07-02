package v3

import (
	"gomessage/models"
	"gomessage/pkg/utils"
	"gomessage/services/core/v1"
	"gomessage/services/format"
)

type ClientActionWechatRobot struct {
	Client *models.Client
}

func (c *ClientActionWechatRobot) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	if isMerge {
		msg := v1.MessageJoint(contentList, utils.VarWechatRobot)
		data := format.PackWechatRobotMessage(client.ExtendWechatRobot.RobotKeyword, msg)
		msgList = append(msgList, data)
	} else {
		for _, msg := range contentList {
			data := format.PackWechatRobotMessage(client.ExtendWechatRobot.RobotKeyword, msg)
			msgList = append(msgList, data)
		}
	}
	return msgList
}

func (c *ClientActionWechatRobot) PushMessages(messages []any) {
	url := v1.RobotRandomUrl(c.Client.ExtendWechatRobot.RobotUrlRandomList)
	for _, msg := range messages {
		v1.Push(msg, url)
	}
}
