package send

import (
	"encoding/json"
	"gomessage/apps/controllers/clients"
	"gomessage/apps/models"
)

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
