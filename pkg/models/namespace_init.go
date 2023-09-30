package models

import (
	"gomessage/pkg/utils/database"
	"gomessage/pkg/utils/log/loggers"
)

// InitDefaultNamespace 初始化default命名空间，在main函数中被调用，全局只被调用一次，创建默认的Namespace（并不是多通道复用，因此就不用设计传参了）
func InitDefaultNamespace() {
	ns := "default"
	var queryNamespace Namespace
	result := database.DB.Default.Where(&Namespace{Name: "default"}).First(&queryNamespace)
	if result.Error != nil {
		newNamespace := Namespace{
			IsActive:    true,
			Name:        ns,
			Description: "系统自动创建的\"默认通道\"，可通过 /go/message 或 /go/default 接收消息推送。",
		}
		database.DB.Default.Create(&newNamespace)

		//初始化模板
		InitTemplate(ns)

		//初始化变量映射
		InitVarMap(ns)

		loggers.DefaultLogger.Info("创建default命名空间...")
	} else {
		loggers.DefaultLogger.Info("default命名空间已存在...")
	}
}

// InitTemplate 初始化信息（多通道复用）
func InitTemplate(ns string) {
	content := `{{ if eq .N6 "firing" }}

## <font color='#FF0000'>【报警中】服务器{{ .N3 }}</font>

{{ else if eq .N6 "resolved" }}

## <font color='#20B2AA'>【已恢复】服务器{{ .N3 }}</font>

{{ else }}

## 标题：信息通知

{{ end }}

====================

**告警规则**：{{ .N1 }}

**告警级别**：{{ .N2 }}

**主机名称**：{{ .N3 }} 

**主机地址**：{{ .N4 }}

**告警详情**：{{ .N5 }}

**告警状态**：{{ .N6 }}

**触发时间**：{{ .N7 }}

**发送时间**：{{ .N8 }}

**规则详情**：[Prometheus控制台](https://www.baidu.com)

**报警详情**：[Alertmanager控制台](https://www.baidu.com)
`

	defaultTemplate := Template{
		Namespace:       ns,
		TemplateName:    ns,
		TemplateContent: content,
		TemplateIsMerge: false,
	}
	AddTemplate(&defaultTemplate)
}

// InitVarMap 初始化变量映射（多通道复用）
func InitVarMap(ns string) {
	keyValueList := []map[string]string{
		{"N1": "alerts.#.labels.alertname"},
		{"N2": "alerts.#.labels.severity"},
		{"N3": "alerts.#.labels.hostname"},
		{"N4": "alerts.#.labels.ping"},
		{"N5": "alerts.#.annotations.description"},
		{"N6": "status"},
		{"N7": "alerts.#.startsAt"},
		{"N8": "alerts.#.endsAt"},
	}
	UpdateAddVars(ns, keyValueList)
}