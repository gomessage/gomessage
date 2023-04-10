package send

import (
	"encoding/json"
	"fmt"
	"gomessage/apps/controllers/clients"
	"gomessage/apps/models"
	"strings"
)

// MessageJoint 按照不同的客户端类型，选择不同的拼接和聚合形式
func MessageJoint(messageList []string, thisType string) string {
	var msg string
	if thisType == "dingtalk" {
		msg = strings.Join(messageList, "\n * * * \n")

	} else if thisType == "wechat" {
		msg = strings.Join(messageList, "\n · \n")

	} else if thisType == "feishu" {
		msg = strings.Join(messageList, "\n -------------- \n")

	} else {
		fmt.Println("客户端类型错误")
	}
	return msg
}

func FeishuMessageFormat(isRenders bool, isMerge bool, client *models.Client, contentList []string) (string, []any) {
	var msgList []any
	url := RobotRandomUrl(client.ExtendFeishu.RobotUrls)
	if isRenders {
		if isMerge {
			msg := MessageJoint(contentList, "feishu")
			data := clients.PackFeishuMessage(client, msg)
			msgList = append(msgList, data)
		} else {
			for _, msg := range contentList {
				data := clients.PackFeishuMessage(client, msg)
				msgList = append(msgList, data)
			}
		}
		return url, msgList
	} else {
		for _, msg := range contentList {
			var cc map[string]any
			requestByte := []byte(msg)
			json.Unmarshal(requestByte, &cc)
			msgList = append(msgList, cc)
		}
		return url, msgList
	}

}

func DingtalkMessageFormat(isRenders, isMerge bool, client *models.Client, contentList []string) (string, []any) {
	var msgList []any
	url := RobotRandomUrl(client.ExtendDingtalk.RobotUrlInfoList)
	if isRenders {
		if isMerge {
			msg := MessageJoint(contentList, "dingtalk")
			data := clients.PackDingtalkMessage(client.ExtendDingtalk.RobotKeyword, msg)
			msgList = append(msgList, data)
		} else {
			for _, msg := range contentList {
				data := clients.PackDingtalkMessage(client.ExtendDingtalk.RobotKeyword, msg)
				msgList = append(msgList, data)
			}
		}
		return url, msgList
	} else {
		for _, msg := range contentList {
			var cc map[string]any
			requestByte := []byte(msg)
			json.Unmarshal(requestByte, &cc)
			msgList = append(msgList, cc)
		}
		return url, msgList
	}
}
