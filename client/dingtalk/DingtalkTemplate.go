package dingtalk

import "time"

//########################
//结构体：接收Alertmanager推送过来的数据信息
//########################
type Messages struct {
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
	Alerts   []struct {
		Status      string            `json:"status"`
		Labels      map[string]string `json:"labels"`
		Annotations struct {
			Description string `json:"description"`
			RunbookURL  string `json:"runbook_url"`
			Summary     string `json:"summary"`
		} `json:"annotations"`
		StartsAt     time.Time `json:"startsAt"`
		EndsAt       time.Time `json:"endsAt"`
		GeneratorURL string    `json:"generatorURL"`
		Fingerprint  string    `json:"fingerprint"`
	} `json:"alerts"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
}

//########################
//template：用来拼装出来要转发出去的报警信息
//########################
const K8sMessageTemplate = `
{{ if eq .N7 "firing" }}
## <font color='#FF0000'>【报警中】Kubernetes集群</font>
{{ else if eq .N7 "resolved" }}
## <font color='#20B2AA'>【已恢复】Kubernetes集群</font>
{{ end }}

========================

**告警规则**：{{ .N1 }}

**告警级别**：{{ .N2 }}

**告警Pod**：{{ .N3 }} 

**命名空间**：{{ .N4 }}

**虚拟地址**：{{ .N5 }}

**告警详情**：{{ .N6 }}

**告警状态**：{{ .N7 }}

**触发时间**：{{ .N8 }}

**发送时间**：{{ .N9 }}

**规则详情**：[Prometheus控制台]({{ .N10 }})

**报警详情**：[Alertmanager控制台]({{ .N11 }})

`

//########################
//template：用来拼装出来要转发出去的报警信息
//########################
const LinuxMessageTemplate = `
{{ if eq .N7 "firing" }}
## <font color='#FF0000'>【报警中】服务器{{ .N3 }}</font>
{{ else if eq .N7 "resolved" }}
## <font color='#20B2AA'>【已恢复】服务器{{ .N3 }}</font>
{{ end }}

========================

**告警规则**：{{ .N1 }}

**告警级别**：{{ .N2 }}

**主机名称**：{{ .N3 }} 

**主机地址**：{{ .N4 }}

**告警详情**：{{ .N6 }}

**告警状态**：{{ .N7 }}

**触发时间**：{{ .N8 }}

**发送时间**：{{ .N9 }}

**规则详情**：[Prometheus控制台]({{ .N10 }})

**报警详情**：[Alertmanager控制台]({{ .N11 }})

`
