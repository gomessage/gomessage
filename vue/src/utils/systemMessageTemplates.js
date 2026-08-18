const commonAlertTemplate = `{{ if eq .N5 "firing" }}
## 🚨 【报警中】服务器 {{ .N3 }}
{{ else if eq .N5 "resolved" }}
## ✅ 【已恢复】服务器 {{ .N3 }}
{{ else }}
## ℹ️ 信息通知
{{ end }}

**告警规则**：{{ .N1 }}

**告警级别**：{{ .N2 }}

**主机名称**：{{ .N3 }}

**告警详情**：{{ .N4 }}

**告警状态**：{{ .N5 }}

**触发时间**：{{ .N6 }}
{{ if eq .N5 "resolved" }}
**恢复时间**：{{ .N7 }}
{{ end }}
{{ if .N8 }}
**规则详情**：[Prometheus 控制台]({{ .N8 }})
{{ end }}
{{ if .N9 }}
**报警详情**：[Alertmanager 控制台]({{ .N9 }})
{{ end }}`

const originalBuiltInTemplate = `{{ if eq .N5 "firing" }}

## <font color='#FF0000'>【报警中】服务器{{ .N3 }}</font>

{{ else if eq .N5 "resolved" }}

## <font color='#20B2AA'>【已恢复】服务器{{ .N3 }}</font>

{{ else }}

## 标题：信息通知

{{ end }}

====================

**告警规则**：{{ .N1 }}

**告警级别**：{{ .N2 }}

**主机名称**：{{ .N3 }}

**告警详情**：{{ .N4 }}

**告警状态**：{{ .N5 }}

**触发时间**：{{ .N6 }}
{{ if eq .N5 "resolved" }}
**恢复时间**：{{ .N7 }}
{{ end }}
**规则详情**：[Prometheus控制台]({{ .N8 }})

**报警详情**：[Alertmanager控制台]({{ .N9 }})`

export const SYSTEM_MESSAGE_TEMPLATES = [
  {
    id: 'original-built-in',
    platform: '通用（原始）',
    name: '原始内置模板',
    description: '完整保留最初内置模板及 #FF0000、#20B2AA 十六进制颜色。各通道的 Markdown 兼容性不同，飞书请优先选用飞书专用模板。',
    content: originalBuiltInTemplate
  },
  {
    id: 'common-alertmanager',
    platform: '通用',
    name: 'Alertmanager 基础告警',
    description: '仅使用常见 Markdown 语法，适合同一命名空间同时配置多种通道。',
    content: commonAlertTemplate
  },
  {
    id: 'dingtalk-alertmanager',
    platform: '钉钉机器人',
    name: 'Alertmanager 告警',
    description: '使用钉钉机器人支持的标题、粗体、表情和链接，不包含平台专属颜色标签。',
    content: `{{ if eq .N5 "firing" }}
## 🚨 【钉钉告警】{{ .N1 }}
{{ else if eq .N5 "resolved" }}
## ✅ 【钉钉恢复】{{ .N1 }}
{{ else }}
## ℹ️ 【钉钉通知】{{ .N1 }}
{{ end }}

**告警级别**：{{ .N2 }}

**故障主机**：{{ .N3 }}

**告警内容**：{{ .N4 }}

**当前状态**：{{ .N5 }}

**触发时间**：{{ .N6 }}
{{ if eq .N5 "resolved" }}
**恢复时间**：{{ .N7 }}
{{ end }}
{{ if .N8 }}
[查看规则详情]({{ .N8 }})
{{ end }}
{{ if .N9 }}
[打开 Alertmanager]({{ .N9 }})
{{ end }}`
  },
  {
    id: 'feishu-alertmanager',
    platform: '飞书机器人',
    name: 'Card JSON 2.0 告警',
    description: '使用飞书卡片 Markdown 支持的颜色名，报警为红色，恢复为绿色。',
    content: `{{ if eq .N5 "firing" }}
## <font color='red'>🚨【报警中】服务器 {{ .N3 }}</font>
{{ else if eq .N5 "resolved" }}
## <font color='green'>✅【已恢复】服务器 {{ .N3 }}</font>
{{ else }}
## <font color='grey'>ℹ️ 信息通知</font>
{{ end }}

**告警规则**：{{ .N1 }}

**告警级别**：{{ .N2 }}

**主机名称**：{{ .N3 }}

**告警详情**：{{ .N4 }}

**告警状态**：{{ .N5 }}

**触发时间**：{{ .N6 }}
{{ if eq .N5 "resolved" }}
**恢复时间**：{{ .N7 }}
{{ end }}
{{ if .N8 }}
**规则详情**：[Prometheus 控制台]({{ .N8 }})
{{ end }}
{{ if .N9 }}
**报警详情**：[Alertmanager 控制台]({{ .N9 }})
{{ end }}`
  },
  {
    id: 'wechat-robot-alertmanager',
    platform: '企业微信机器人',
    name: 'Alertmanager 告警',
    description: '使用企业微信 Markdown 的 warning、info 和 comment 颜色名。',
    content: `{{ if eq .N5 "firing" }}
## <font color="warning">🚨【报警中】{{ .N1 }}</font>
{{ else if eq .N5 "resolved" }}
## <font color="info">✅【已恢复】{{ .N1 }}</font>
{{ else }}
## <font color="comment">ℹ️ 信息通知</font>
{{ end }}

> 告警级别：<font color="warning">{{ .N2 }}</font>
> 故障主机：{{ .N3 }}
> 告警内容：{{ .N4 }}
> 当前状态：{{ .N5 }}
> 触发时间：{{ .N6 }}
{{ if eq .N5 "resolved" }}
> 恢复时间：{{ .N7 }}
{{ end }}
{{ if .N8 }}
[查看规则详情]({{ .N8 }})
{{ end }}
{{ if .N9 }}
[打开 Alertmanager]({{ .N9 }})
{{ end }}`
  },
  {
    id: 'wechat-app-alertmanager',
    platform: '企业微信应用号',
    name: 'Alertmanager 应用通知',
    description: '适合应用号单独推送，使用企业微信 Markdown 颜色名强调状态和关键信息。',
    content: `{{ if eq .N5 "firing" }}
## <font color="warning">【系统告警】{{ .N1 }}</font>
{{ else if eq .N5 "resolved" }}
## <font color="info">【系统恢复】{{ .N1 }}</font>
{{ else }}
## <font color="comment">【系统通知】{{ .N1 }}</font>
{{ end }}

**服务器**：{{ .N3 }}

**级别**：<font color="warning">{{ .N2 }}</font>

**详情**：{{ .N4 }}

**状态**：{{ .N5 }}

**触发时间**：{{ .N6 }}
{{ if eq .N5 "resolved" }}
**恢复时间**：{{ .N7 }}
{{ end }}
{{ if .N9 }}
[查看告警中心]({{ .N9 }})
{{ end }}`
  }
]

export {commonAlertTemplate}
