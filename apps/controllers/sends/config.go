package sends

import (
	"gomessage/apps/models"
	"math/rand"
	"time"
)

type UserConfig struct {
	Namespace       string
	VariablesMap    []map[string]string
	MessageTemplate string
	MessageMerge    bool
	ActiveClient    []models.Client
}

// GetUserConfig 获取用户在图形界面上设置的各种参数
func GetUserConfig(ns string) UserConfig {
	c := UserConfig{Namespace: ns}

	//获取变量映射（根据命名空间获取）
	listVariables, err := models.ListVariables(c.Namespace)
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
	template, err := models.GetTemplateByNamespace(c.Namespace)
	if err != nil {
		panic(err)
	}
	c.MessageTemplate = template.TemplateContent
	c.MessageMerge = template.TemplateIsMerge

	//获取客户端相关（根据命名空间获取）
	c.ActiveClient, _ = models.GetActiveClient(c.Namespace)
	return c
}

// RobotRandomUrl 随机获取一个机器人地址（通用方法：可以同时被钉钉和飞书使用）
func RobotRandomUrl(rList []string) string {
	rand.Seed(time.Now().Unix())
	i := rand.Int() % len(rList)
	return rList[i]
}