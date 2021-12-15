package alertmanager

import (
	"GoMessage/client/dingtalk"
	web2 "GoMessage/controllers/web2"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

// 接收Alertmanager推送的报警信息
type K8sControllers struct {
	beego.Controller
}

// @Title 推送数据
// @Description 推送数据到钉钉，数据来自于Alertmanager的webhook推送
// @Success 200 ok
// @router /dingding/k8s [post]
func (this *K8sControllers) Post() {

	msg := dingtalk.Messages{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &msg)
	if err != nil {
		fmt.Errorf("错误：%v", err)
	}

	web2.TmpJsonData.JsonData = msg
	web2.TmpJsonData.UpdateTime = time.Now()

	//从request推送过来的数据中，拿到需要转发出去的数据，然后实例化为一个结构体对象，传递给下文使用
	messageData := dingtalk.MessageStruct{
		N1:  msg.Alerts[0].Labels["alertname"],
		N2:  msg.Alerts[0].Labels["severity"],
		N3:  msg.Alerts[0].Labels["pod"],
		N4:  msg.Alerts[0].Labels["namespace"],
		N5:  msg.Alerts[0].Labels["instance"],
		N6:  msg.Alerts[0].Annotations.Description,
		N7:  msg.Alerts[0].Status,
		N8:  msg.Alerts[0].StartsAt.Local().Format("2006-01-02 15:04:05"),
		N9:  msg.Alerts[0].EndsAt.Local().Format("2006-01-02 15:04:05"),
		N10: "http://192.168.35.18:9090/alerts",
		N11: "http://192.168.35.18:9093/#/alerts",
	}
	d := dingtalk.NewDingTalk(dingtalk.GetDingtalkRobotURl(), messageData, "k8s")
	result := d.SendMessage() //调用信息发送函数
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = struct {
		Result interface{} `json:"dingding_response"`
	}{Result: result}
	this.ServeJSON()
}
