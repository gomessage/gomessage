package prometheus

import "time"

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

//推送过来的消息（结构体），用于解析json类型的数据
//type MessagesJoins struct {
//	Receiver string `json:"receiver"`
//	Status   string `json:"status"`
//	Alerts   []struct {
//		Status string `json:"status"`
//		Labels map[string]string `json:"labels"`
//		Annotations struct {
//			Description string `json:"description"`
//			RunbookURL  string `json:"runbook_url"`
//			Summary     string `json:"summary"`
//		} `json:"annotations,omitempty"`
//		StartsAt     time.Time `json:"startsAt"`
//		EndsAt       time.Time `json:"endsAt"`
//		GeneratorURL string    `json:"generatorURL"`
//		Fingerprint  string    `json:"fingerprint"`
//	} `json:"alerts"`
//	GroupLabels struct {
//	} `json:"groupLabels,omitempty"`
//	CommonLabels struct {
//		Prometheus string `json:"prometheus"`
//	} `json:"commonLabels"`
//	CommonAnnotations struct {
//	} `json:"commonAnnotations"`
//	ExternalURL     string `json:"externalURL"`
//	Version         string `json:"version"`
//	GroupKey        string `json:"groupKey"`
//	TruncatedAlerts int    `json:"truncatedAlerts"`
//}

//运维测试机器人的URL地址
//var YunweiRobotList = [...]string{
//	"https://oapi.prometheus.com/robot/send?access_token=17cdf0818a83569427241b7019a3fe556824f1849ff9f8806bc0f83339a45c5e",
//	"https://oapi.prometheus.com/robot/send?access_token=678759e5a116f4d405083f7130f1f2c8654edd4d98225cba8dca1e72def0788f",
//	"https://oapi.prometheus.com/robot/send?access_token=759b795bc530a6e19c3746f371b11a3beafb82a0f91f719fe7ad312bdb3b0156",
//	"https://oapi.prometheus.com/robot/send?access_token=2fbc42d100027483133d2e314aaf45d9c31485e4468f391f51ee8b9ed6c2b7a9",
//	"https://oapi.prometheus.com/robot/send?access_token=40667f16f3ef04c79a3a22467950095abd99f8be54e7a995e81a1dc7d2ec6eec",
//	"https://oapi.prometheus.com/robot/send?access_token=abe8be7799f3061a3e1b253e21fad89f3f13628d9ecce2e98a1cbbd87b96ca63",
//	"https://oapi.prometheus.com/robot/send?access_token=7c94fd96db657713c5e0406202c9395d13dacdc2307b180a4f4ecc3cc2140f7f",
//	"https://oapi.prometheus.com/robot/send?access_token=da2ed46d6a6f538c2ee47790580435dfe70cfa8e9adb94565b54bf6b917ce302",
//}
