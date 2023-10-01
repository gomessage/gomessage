package v1

import (
	"fmt"
	"gomessage/pkg/models"
	"gomessage/pkg/utils"
)

// RendersRequestData 渲染数据
//func RendersRequestData(isRenders bool, thisNamespaceUserConfig NamespaceUserConfig, requestByte []byte) []string {
//	var contentList []string
//	//判断是否渲染
//	if isRenders {
//		//得到变量映射
//		analysisDataList := AnalysisData(thisNamespaceUserConfig.VariablesMap, requestByte)
//		//得到内容体
//		contentList = CompleteMessage(thisNamespaceUserConfig.MessageTemplate, analysisDataList)
//	} else {
//		//不渲染数据
//		contentList = append(contentList, string(requestByte))
//	}
//	return contentList
//}

type ReadyClient struct {
	Id      int
	Name    string
	Type    string
	IsMerge bool
	Url     string
	Data    []any
}

func AssembledMessage(isRenders bool, thisNamespaceUserConfig NamespaceUserConfig, contentList []string) []ReadyClient {
	var clientDataList []ReadyClient

	//遍历客户端
	for _, client := range thisNamespaceUserConfig.ActiveClient {
		clientInfo, _ := models.GetClientById(client.ID)

		cd := ReadyClient{
			Id:      clientInfo.ID,
			Name:    clientInfo.ClientName,
			IsMerge: thisNamespaceUserConfig.IsMerge,
			Type:    clientInfo.ClientType,
		}

		switch clientInfo.ClientType {
		case utils.VarFeishu:
			cd.Url, cd.Data = FeishuMessageFormat(isRenders, cd.IsMerge, clientInfo, contentList)

		case utils.VarDingtalk:
			cd.Url, cd.Data = DingtalkMessageFormat(isRenders, cd.IsMerge, clientInfo, contentList)

		case utils.VarWechatRobot:
			fmt.Println("后面补充")
		}
		clientDataList = append(clientDataList, cd)
	}
	return clientDataList
}
