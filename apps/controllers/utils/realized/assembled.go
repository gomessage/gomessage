package realized

import (
	"encoding/json"
	"fmt"
	"gomessage/apps/controllers/clients"
	"gomessage/apps/controllers/send"
	"gomessage/apps/controllers/utils/base"
	"gomessage/apps/models"
)

// GeneralMessageAssembled 通用消息体封装
type GeneralMessageAssembled struct {
	base.Assembled
}

func (d *GeneralMessageAssembled) AssembledData(isRenders, isMerge bool, client *models.Client, contentList []string) (string, []any) {
	fmt.Println("消息体原封不动向后传递，不做任何变动...")

	var msgList []any
	for _, msg := range contentList {
		var cc map[string]any
		requestByte := []byte(msg)
		json.Unmarshal(requestByte, &cc)
		msgList = append(msgList, cc)
	}

	return "", msgList
}

// DingtalkMessageAssembled 钉钉消息体封装
type DingtalkMessageAssembled struct {
	base.Assembled
}

func (d *DingtalkMessageAssembled) AssembledData(isRenders, isMerge bool, client *models.Client, contentList []string) (string, []any) {
	fmt.Println("组装钉钉需要的消息体...")
	var msgList []any
	url := send.RobotRandomUrl(client.ExtendDingtalk.RobotUrlInfoList)

	if isRenders {
		if isMerge {
			msg := send.MessageJoint(contentList, "dingtalk")
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

// FeishuMessageAssembled 飞书消息体封装
type FeishuMessageAssembled struct {
	base.Assembled
}

func (d *FeishuMessageAssembled) AssembledData(isRenders, isMerge bool, client *models.Client, contentList []string) (string, []any) {
	fmt.Println("组装飞书需要的消息体...")
	var msgList []any
	url := send.RobotRandomUrl(client.ExtendFeishu.RobotUrls)

	if isRenders {
		if isMerge {
			msg := send.MessageJoint(contentList, "feishu")
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

// WechatMessageAssembled 企业微信消息体封装
type WechatMessageAssembled struct {
	base.Assembled
}

func (d *WechatMessageAssembled) AssembledData(isRenders, isMerge bool, client *models.Client, contentList []string) (string, []any) {
	fmt.Println("组装企业微信需要的消息体...")
	var msgList []any
	url := ""

	if isRenders {
		if isMerge {
			msg := send.MessageJoint(contentList, "wechat")
			msgList = append(msgList, msg)
		} else {
			for _, msg := range contentList {
				msgList = append(msgList, msg)
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

// EmailMessageAssembled 邮箱消息体封装
type EmailMessageAssembled struct {
	base.Assembled
}

func (d *EmailMessageAssembled) AssembledData(isRenders, isMerge bool, client *models.Client, contentList []string) (string, []any) {
	fmt.Println("组装为Email需要的消息体...")
	return "", nil
}
