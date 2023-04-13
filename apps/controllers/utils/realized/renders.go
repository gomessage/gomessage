package realized

import (
	"fmt"
	"gomessage/apps/controllers/send"
	"gomessage/apps/controllers/utils/base"
)

// GetRendersResult 开启渲染的实现
type GetRendersResult struct {
	base.Renders
	Rds bool
}

func (d *GetRendersResult) RendersData(thisNamespaceUserConfig send.NamespaceUserConfig, requestByte []byte) []string {
	fmt.Println("渲染为钉钉需要的内容体...")
	var contentList []string
	//判断是否需要渲染
	if d.Rds {
		//得到变量映射
		analysisDataList := send.AnalysisData(thisNamespaceUserConfig.VariablesMap, requestByte)
		//得到内容体
		contentList = send.CompleteMessage(thisNamespaceUserConfig.MsgTemplate, analysisDataList)
	} else {
		contentList = append(contentList, string(requestByte))
	}
	return contentList
}
