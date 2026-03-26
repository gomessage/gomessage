package v3

import (
	"gomessage/pkg/models"
	v12 "gomessage/pkg/services/core/v1"
	"gomessage/pkg/services/format"
	"gomessage/pkg/utils"
)

type ClientActionWechatRobot struct {
	Client *models.Client
}

func (c *ClientActionWechatRobot) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	if isMerge {
		msg := v12.MessageJoint(contentList, utils.VarWechatRobot)
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

func (c *ClientActionWechatRobot) PushMessages(messages []any) (successCount int, failureCount int) {
	url := v12.RobotRandomUrl(c.Client.ExtendWechatRobot.RobotUrlRandomList)
	for _, msg := range messages {
		if err := v12.Push(msg, url); err != nil {
			failureCount++
		} else {
			successCount++
		}
	}
	return
}
