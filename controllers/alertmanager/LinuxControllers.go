package alertmanager

import (
	"GoMessage/client/dingtalk"
	web2 "GoMessage/controllers/web2"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

// 转发来自于Prometheus的数据
type LinuxControllers struct {
	beego.Controller
}

// @Title 推送数据
// @Description 推送数据到钉钉，数据来自于Alertmanager的webhook推送
// @Success 200 ok
// @router /dingding/linux [post]
func (this *LinuxControllers) Post() {
	//接受来自于Alertmanager的webhook推送，解析后的json数据存放在结构体msg中
	msg := dingtalk.Messages{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &msg)
	if err != nil {
		fmt.Errorf("错误：%v", err)
	}

	//绑定数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &web2.CacheData.JsonData)
	web2.CacheData.UpdateTime = time.Now()

	//fmt.Println(msg.Alerts[0].StartsAt)
	//fmt.Println(msg.Alerts[0].StartsAt.Add(8 * time.Hour))
	//fmt.Println(msg.Alerts[0].StartsAt.Local())
	//fmt.Println(msg.Alerts[0].StartsAt.Location())
	//msg.Alerts[0].StartsAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05")
	//msg.Alerts[0].EndsAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05")
	//msg.Alerts[0].StartsAt.Local().Format("2006-01-02 15:04:05")
	//msg.Alerts[0].EndsAt.Local().Format("2006-01-02 15:04:05")

	//不同的template模板需要使用到的数据不同（在这一步完成对占位符的拼装）
	//从request推送过来的数据中，拿到需要转发出去的数据，然后实例化为一个结构体对象，传递给下文使用
	messageData := dingtalk.MessageStruct{
		N1:  msg.Alerts[0].Labels["alertname"],
		N2:  msg.Alerts[0].Labels["severity"],
		N3:  msg.Alerts[0].Labels["hostname"],
		N4:  msg.Alerts[0].Labels["ping"],
		N6:  msg.Alerts[0].Annotations.Description,
		N7:  msg.Alerts[0].Status,
		N8:  msg.Alerts[0].StartsAt.Local().Format("2006-01-02 15:04:05"),
		N9:  msg.Alerts[0].EndsAt.Local().Format("2006-01-02 15:04:05"),
		N10: "http://192.168.35.18:9090/alerts",
		N11: "http://192.168.35.18:9093/#/alerts",
	}
	d := dingtalk.NewDingTalk(dingtalk.GetDingtalkRobotURl(), messageData, "linux")
	result := d.SendMessage() //调用信息发送函数
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = struct {
		Result interface{} `json:"dingding_response"`
	}{Result: result}
	this.ServeJSON()
}
