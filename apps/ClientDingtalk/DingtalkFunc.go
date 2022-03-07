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
		//以下机器人地址，仅仅是为下一个版本的【新功能的增加】而进行的测试预留，这几个机器人都在【接收侧】指定了IP白名单，请大家忽略以下机器人地址~
		"https://oapi.dingtalk.com/robot/send?access_token=17cdf0818a83569427241b7019a3fe556824f1849xxxxxxxxxxxxxxxxxxxxxxx",
		"https://oapi.dingtalk.com/robot/send?access_token=678759e5a116f4d405083f7130f1f2c8654edd4d9xxxxxxxxxxxxxxxxxxxxxxx",
		"https://oapi.dingtalk.com/robot/send?access_token=759b795bc530a6e19c3746f371b11a3beafb82a0fxxxxxxxxxxxxxxxxxxxxxxx",
		"https://oapi.dingtalk.com/robot/send?access_token=2fbc42d100027483133d2e314aaf45d9c31485e44xxxxxxxxxxxxxxxxxxxxxxx",
		"https://oapi.dingtalk.com/robot/send?access_token=40667f16f3ef04c79a3a22467950095abd99f8be5xxxxxxxxxxxxxxxxxxxxxxx",
		"https://oapi.dingtalk.com/robot/send?access_token=abe8be7799f3061a3e1b253e21fad89f3f13628d9xxxxxxxxxxxxxxxxxxxxxxx",
		"https://oapi.dingtalk.com/robot/send?access_token=7c94fd96db657713c5e0406202c9395d13dacdc23xxxxxxxxxxxxxxxxxxxxxxx",
		"https://oapi.dingtalk.com/robot/send?access_token=da2ed46d6a6f538c2ee47790580435dfe70cfa8e9xxxxxxxxxxxxxxxxxxxxxxx",
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
