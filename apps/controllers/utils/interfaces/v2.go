package interfaces

import "gomessage/apps/models"

//var RegistrationCenter = map[string]ClientAction{}
//
//func init() {
//	dingtalk := ClientActionDingtalk{}
//	feishu := ClientActionFeishu{}
//
//	RegistrationCenter["dingtalk"] = &dingtalk
//	RegistrationCenter["feishu"] = &feishu
//}

type ClientAction interface {
	RendersMessages(client *models.Client, isMerge bool, contentList []string) []any
	PushMessages(messages []any)
}
