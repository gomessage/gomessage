package crontab

import (
	"gomessage/pkg/models"
	v3 "gomessage/pkg/services/core/v3"
	"gomessage/pkg/utils"
)

func genJobFunc(ct models.Crontab) func() {
	clients := []*models.Client{}

	for _, nsStr := range ct.CrontabNamespace {
		acList, err := models.GetActiveClient(nsStr)
		if err != nil {
			// TODO: error
			continue
		}
		for i := range acList {
			clients = append(clients, &acList[i])
		}
	}

	return func() {
		for _, c := range clients {
			push(c, ct.CrontabContent)
		}
	}
}

func push(clientInfo *models.Client, content string) {
	var clientAction v3.ClientAction

	switch clientInfo.ClientType {
	default:
		return
	case utils.VarDingtalk:
		clientAction = &v3.ClientActionDingtalk{Client: clientInfo}
	case utils.VarFeishu:
		clientAction = &v3.ClientActionFeishu{Client: clientInfo}
	case utils.VarWechatApplication:
		clientAction = &v3.ClientActionWechatApplication{
			CorpId:      clientInfo.ExtendWechatApplication.CorpId,
			AgentId:     clientInfo.ExtendWechatApplication.AgentId,
			AgentSecret: clientInfo.ExtendWechatApplication.Secret,
			Touser:      clientInfo.ExtendWechatApplication.Touser,
		}
	case utils.VarWechatRobot:
		clientAction = &v3.ClientActionWechatRobot{Client: clientInfo}
	}

	messages := clientAction.RendersMessages(clientInfo, false, []string{content})

	clientAction.PushMessages(messages)
}
