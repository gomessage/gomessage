package v1

import (
	"encoding/json"
	"gomessage/pkg/models"
	"gomessage/pkg/services/format"
	"gomessage/pkg/utils"
)

func FeishuMessageFormat(isRenders bool, isMerge bool, client *models.Client, contentList []string) (string, []any) {
	var msgList []any
	url := RobotRandomUrl(client.ExtendFeishu.RobotUrlRandomList)
	if isRenders {
		if isMerge {
			msg := MessageJoint(contentList, utils.VarFeishu)
			data := format.PackFeishuMessage(client, msg)
			msgList = append(msgList, data)
		} else {
			for _, msg := range contentList {
				data := format.PackFeishuMessage(client, msg)
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
	url := RobotRandomUrl(client.ExtendDingtalk.RobotUrlRandomList)
	if isRenders {
		if isMerge {
			msg := MessageJoint(contentList, utils.VarDingtalk)
			data := format.PackDingtalkMessage(client.ExtendDingtalk.RobotKeyword, msg, false, nil)
			msgList = append(msgList, data)
		} else {
			for _, msg := range contentList {
				data := format.PackDingtalkMessage(client.ExtendDingtalk.RobotKeyword, msg, false, nil)
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
