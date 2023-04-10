package send

import (
	"fmt"
	"gomessage/apps/models"
	"gomessage/utils/log/loggers"
)

// GetNs 获取通道信息
func GetNs(namespace string) *models.Namespace {
	if namespace == "message" {
		namespace = "default"
	}
	nsInfo, _ := models.GetNamespaceByName(namespace)
	loggers.DefaultLogger.Info("消息发送至" + namespace + "通道")
	return nsInfo
}

// RendersRequestData 渲染数据
func RendersRequestData(isRenders bool, thisNamespaceUserConfig ThisNamespaceUserConfig, requestByte []byte) []string {
	var contentList []string
	//判断是否渲染
	if isRenders {
		//得到变量映射
		analysisDataList := AnalysisData(thisNamespaceUserConfig.VariablesMap, requestByte)
		//得到内容体
		contentList = CompleteMessage(thisNamespaceUserConfig.MsgTemplate, analysisDataList)
	} else {
		//不渲染数据
		contentList = append(contentList, string(requestByte))
	}
	return contentList
}

type ReadyClient struct {
	Id      int
	Name    string
	Type    string
	IsMerge bool
	Url     string
	Data    []any
}

func BuilderClient(isRenders bool, thisNamespaceUserConfig ThisNamespaceUserConfig, contentList []string) []ReadyClient {
	var clientDataList []ReadyClient

	//遍历客户端
	for _, client := range thisNamespaceUserConfig.ActiveClient {
		clientInfo, _ := models.GetClientById(client.ID)

		cd := ReadyClient{
			Id:      clientInfo.ID,
			Name:    clientInfo.ClientName,
			IsMerge: thisNamespaceUserConfig.MsgMerge,
			Type:    clientInfo.ClientType,
		}

		switch clientInfo.ClientType {
		case "feishu":
			cd.Url, cd.Data = FeishuMessageFormat(isRenders, cd.IsMerge, clientInfo, contentList)

		case "dingtalk":
			cd.Url, cd.Data = DingtalkMessageFormat(isRenders, cd.IsMerge, clientInfo, contentList)

		case "wechat":
			fmt.Println("后面补充")
		}
		clientDataList = append(clientDataList, cd)
	}
	return clientDataList
}
