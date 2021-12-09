package controllers

import (
	"GoMessage/client/dingtalk"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

//随机返回出去一个钉钉机器人的URL
func getRobotURl() string {
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
func printJson(msg struct{}) {
	bs, _ := json.Marshal(msg)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Printf("msg=%+v\n", out.String())
}

// 客户端：如果想发送到不同的客户端请调用对应的API
type DingtalkControllers struct {
	beego.Controller
}

// @router /dingtalk [post]
func (this *DingtalkControllers) Post() {
	msg := dingtalk.Messages{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &msg)
	if err != nil {
		fmt.Errorf("错误：%v", err)
	}
	//从request推送过来的数据中，拿到需要转发出去的数据，然后实例化为一个结构体对象，传递给下文使用
	messageData := dingtalk.MessageStruct{
		N1:  msg.Alerts[0].Labels["alertname"],
		N2:  msg.Alerts[0].Labels["severity"],
		N3:  "~",
		N4:  "~",
		N5:  "~",
		N6:  msg.Alerts[0].Annotations.Description,
		N7:  msg.Alerts[0].Status,
		N8:  msg.Alerts[0].StartsAt.String(),
		N9:  msg.Alerts[0].EndsAt.String(),
		N10: "www.baidu.com",
		N11: "www.baidu.com",
	}
	d := dingtalk.NewDingTalk(getRobotURl(), messageData)
	result := d.SendMessage() //调用信息发送函数
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = struct {
		Result dingtalk.ResponseData `json:"dingding_response"`
	}{Result: result}
	this.ServeJSON()
}
