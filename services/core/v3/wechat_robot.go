package v3

import (
	"gomessage/models"
	"gomessage/services/clientFormat"
	v12 "gomessage/services/core/v1"
)

type ClientActionWechatRobot struct {
	Client *models.Client
}

func (c *ClientActionWechatRobot) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	if isMerge {
		msg := v12.MessageJoint(contentList, "wechat_robot")
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
	url := v12.RobotRandomUrl(c.Client.ExtendWechatRobot.RobotUrlInfoList)
	for _, msg := range messages {
		v12.Push(msg, url)
	}
}
