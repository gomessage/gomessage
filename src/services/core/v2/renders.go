package v2

import (
	"fmt"
	v12 "gomessage/src/services/core/v1"
)

// GetRendersResult 渲染数据
type GetRendersResult struct {
	Renders
	Rds bool
}

// RendersData 通用渲染方法：如果需要渲染则映射变量和加载内容模板，如果不需要渲染则把原始信息组装为指定格式，然后返回给调用方
func (d *GetRendersResult) RendersData(thisNamespaceUserConfig v12.NamespaceUserConfig, requestByte []byte) []string {
	fmt.Println("渲染为钉钉需要的内容体...")
	var contentList []string
	//判断是否需要渲染
	if d.Rds {
		//得到变量映射
		analysisDataList := v12.AnalysisData(thisNamespaceUserConfig.VariablesMap, requestByte)
		//得到内容体
		contentList = v12.CompleteMessage(thisNamespaceUserConfig.MessageTemplate, analysisDataList)
	} else {
		contentList = append(contentList, string(requestByte))
	}
	return contentList
}
