package v1

import (
	"fmt"
	"gomessage/pkg/models"
	"gomessage/pkg/utils/log/loggers"
	"math/rand"
	"time"
)

type NamespaceUserConfig struct {
	Namespace       string              //命名空间名称
	VariablesMap    []map[string]string //用户变量映射
	MessageTemplate string              //信息模板
	IsMerge         bool                //是否合并
	ActiveClient    []models.Client     //激活的客户端
}

// GetNamespaceUserConfig 获取用户在图形界面上设置的各种参数
func GetNamespaceUserConfig(ns string, IsRenders bool) (NamespaceUserConfig, error) {
	c := NamespaceUserConfig{Namespace: ns}

	activeClient, err := models.GetActiveClient(c.Namespace)
	if err != nil {
		loggers.DefaultLogger.Errorln("获取激活客户端失败：", err)
		return c, err
	}
	c.ActiveClient = activeClient

	//判断是否要启用渲染功能
	if IsRenders {

		//获取变量映射（根据命名空间获取）
		listVariables, err := models.ListVariables(c.Namespace)
		if err != nil {
			loggers.DefaultLogger.Errorln("获取变量映射失败：", err)
			return c, fmt.Errorf("获取变量映射失败：%w", err)
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
			loggers.DefaultLogger.Errorln("获取模板失败：", err)
			return c, fmt.Errorf("获取模板失败：%w", err)
		}
		c.MessageTemplate = template.TemplateContent
		c.IsMerge = template.TemplateIsMerge
	}

	return c, nil
}

// RobotRandomUrl 随机获取一个机器人地址（通用方法：可以同时被钉钉和飞书使用）
func RobotRandomUrl(rList []string) string {
	rand.Seed(time.Now().Unix())
	i := rand.Int() % len(rList)
	return rList[i]
}

// GetNsInfo 获取通道信息
func GetNsInfo(namespace string) *models.Namespace {
	if namespace == "message" {
		namespace = "default"
	}
	nsInfo, err := models.GetNamespaceByName(namespace)
	if err != nil {
		loggers.DefaultLogger.Errorln("获取命名空间失败：", err)
		return &models.Namespace{Name: namespace}
	}
	return nsInfo
}
