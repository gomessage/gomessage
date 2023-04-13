package realized

import (
	"fmt"
	"gomessage/apps/controllers/clients"
	"gomessage/apps/controllers/send"
	"gomessage/apps/controllers/utils/base"
	"gomessage/apps/models"
)

// DingtalkMessageAssembled 钉钉消息体封装
type DingtalkMessageAssembled struct {
	base.Assembled
}

func (d *DingtalkMessageAssembled) AssembledData(isMerge bool, client *models.Client, contentList []string) (string, []any) {
	fmt.Println("组装为钉钉需要的消息体...")
	var msgList []any
	url := send.RobotRandomUrl(client.ExtendDingtalk.RobotUrlInfoList)
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
}

// FeishuMessageAssembled 飞书消息体封装
type FeishuMessageAssembled struct {
	base.Assembled
}

func (d *FeishuMessageAssembled) AssembledData(isMerge bool, client *models.Client, contentList []string) (string, []any) {
	fmt.Println("组装为飞书需要的消息体...")
	return "", nil
}

// WechatMessageAssembled 微信消息体封装
type WechatMessageAssembled struct {
	base.Assembled
}

func (d *WechatMessageAssembled) AssembledData(isMerge bool, client *models.Client, contentList []string) (string, []any) {
	fmt.Println("组装为企业微信需要的消息体...")
	return "", nil
}

// EmailMessageAssembled 邮箱消息体封装
type EmailMessageAssembled struct {
	base.Assembled
}

func (d *EmailMessageAssembled) AssembledData(isMerge bool, client *models.Client, contentList []string) (string, []any) {
	fmt.Println("组装为Email需要的消息体...")
	return "", nil
}

// GeneralMessageAssembled 通用消息体封装
type GeneralMessageAssembled struct {
	base.Assembled
}

func (d *GeneralMessageAssembled) AssembledData(isMerge bool, client *models.Client, contentList []string) (string, []any) {
	fmt.Println("消息体原封不动，不做任何操作...")

	return "", nil
}
