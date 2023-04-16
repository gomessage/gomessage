package v3

import (
	"gomessage/apps/controllers/clientFormat"
	"gomessage/apps/controllers/core/v1"
	"gomessage/apps/models"
)

type ClientActionWechatRobot struct {
	Client *models.Client
}

func (c *ClientActionWechatRobot) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	if isMerge {
		msg := v1.MessageJoint(contentList, "wechat_robot")
		data := clientFormat.PackWechatRobotMessage(client.ExtendWechatRobot.RobotKeyword, msg)
		msgList = append(msgList, data)
	} else {
		for _, msg := range contentList {
			data := clientFormat.PackWechatRobotMessage(client.ExtendWechatRobot.RobotKeyword, msg)
			msgList = append(msgList, data)
		}
	}
	return msgList
}

func (c *ClientActionWechatRobot) PushMessages(messages []any) {
	url := v1.RobotRandomUrl(c.Client.ExtendWechatRobot.RobotUrlInfoList)
	for _, msg := range messages {
		v1.Push(msg, url)
	}
}
