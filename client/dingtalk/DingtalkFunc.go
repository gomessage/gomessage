package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type MessageStruct struct {
	N1  string
	N2  string
	N3  string
	N4  string
	N5  string
	N6  string
	N7  string
	N8  string
	N9  string
	N10 string
	N11 string
}

//随机返回出去一个钉钉机器人的URL
func GetDingtalkRobotURl() string {
	var yunweiRobotUrl = [...]string{
		"https://oapi.dingtalk.com/robot/send?access_token=17cdf0818a83569427241b7019a3fe556824f1849ff9f8806bc0f83339a45c5e",
		"https://oapi.dingtalk.com/robot/send?access_token=678759e5a116f4d405083f7130f1f2c8654edd4d98225cba8dca1e72def0788f",
		"https://oapi.dingtalk.com/robot/send?access_token=759b795bc530a6e19c3746f371b11a3beafb82a0f91f719fe7ad312bdb3b0156",
		"https://oapi.dingtalk.com/robot/send?access_token=2fbc42d100027483133d2e314aaf45d9c31485e4468f391f51ee8b9ed6c2b7a9",
		"https://oapi.dingtalk.com/robot/send?access_token=40667f16f3ef04c79a3a22467950095abd99f8be54e7a995e81a1dc7d2ec6eec",
		"https://oapi.dingtalk.com/robot/send?access_token=abe8be7799f3061a3e1b253e21fad89f3f13628d9ecce2e98a1cbbd87b96ca63",
		"https://oapi.dingtalk.com/robot/send?access_token=7c94fd96db657713c5e0406202c9395d13dacdc2307b180a4f4ecc3cc2140f7f",
		"https://oapi.dingtalk.com/robot/send?access_token=da2ed46d6a6f538c2ee47790580435dfe70cfa8e9adb94565b54bf6b917ce302",
	}
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(yunweiRobotUrl)
	return yunweiRobotUrl[n]
}

//把结构体转换成json结构，并在控制台打印出来
func PrintJson(msg struct{}) {
	bs, _ := json.Marshal(msg)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Printf("msg=%+v\n", out.String())
}

//实例化DingTalk对象
func NewDingTalk(url string, messageData MessageStruct, alertType string) DingTalk {

	p := DingTalk{}

	if alertType == "k8s" {
		p.RobotUrl = url
		p.MessageStruct = messageData
		p.Template = K8sMessageTemplate
	}
	if alertType == "linux" {
		p.RobotUrl = url
		p.MessageStruct = messageData
		p.Template = LinuxMessageTemplate
	}
	return p
}
