package v1

import (
	models2 "gomessage/src/models"
	"math/rand"
	"time"
)

type NamespaceUserConfig struct {
	Namespace       string              //命名空间名称
	VariablesMap    []map[string]string //用户变量映射
	MessageTemplate string              //信息模板
	IsMerge         bool                //是否合并
	ActiveClient    []models2.Client    //激活的客户端
}

// GetNamespaceUserConfig 获取用户在图形界面上设置的各种参数
func GetNamespaceUserConfig(ns string, IsRenders bool) NamespaceUserConfig {
	c := NamespaceUserConfig{Namespace: ns}

	//获取客户端相关（根据命名空间获取）
	c.ActiveClient, _ = models2.GetActiveClient(c.Namespace)

	//判断是否要启用渲染功能
	if IsRenders {

		//获取变量映射（根据命名空间获取）
		listVariables, err := models2.ListVariables(c.Namespace)
		if err != nil {
			panic(err)
		}
		var varList []map[string]string
		for _, value := range *listVariables {
			tmpMap := make(map[string]string)
			tmpMap[value.Key] = value.Value
			varList = append(varList, tmpMap)
		}
		c.VariablesMap = varList

		//获取消息模板（根据命名空间获取）
		template, err := models2.GetTemplateByNamespace(c.Namespace)
		if err != nil {
			panic(err)
		}
		c.MessageTemplate = template.TemplateContent
		c.IsMerge = template.TemplateIsMerge
	}

	return c
}

// RobotRandomUrl 随机获取一个机器人地址（通用方法：可以同时被钉钉和飞书使用）
func RobotRandomUrl(rList []string) string {
	rand.Seed(time.Now().Unix())
	i := rand.Int() % len(rList)
	return rList[i]
}

// GetNsInfo 获取通道信息
func GetNsInfo(namespace string) *models2.Namespace {
	if namespace == "message" {
		namespace = "default"
	}
	nsInfo, _ := models2.GetNamespaceByName(namespace)
	return nsInfo
}
