package controllers

import (
	"GoMessage/apps/client/dingtalk"
	"GoMessage/apps/prometheus"
	"encoding/json"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

// 转发Prometheus报警
type DingtalkControllers struct {
	beego.Controller
}

// @router / [post]
func (this *DingtalkControllers) Post() {
	msg := prometheus.Messages{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &msg)
	if err != nil {
		fmt.Errorf("错误：%v", err)
	}
	url := "https://oapi.dingtalk.com/robot/send?access_token=678759e5a116f4d405083f7130f1f2c8654edd4d98225cba8dca1e72def0788f"
	template := "MessageTemplate"
	message := dingtalk.MessageStruct{
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

	d := dingtalk.NewDingTalk(url, template, message)
	d.SendMessage()

	//bs, _ := json.Marshal(msg)
	//var out bytes.Buffer
	//json.Indent(&out, bs, "", "\t")
	//fmt.Printf("msg=%+v\n", out.String())

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = struct {
		Message string `json:"message"`
	}{Message: "ok"}
	this.ServeJSON()
}
