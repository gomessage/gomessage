package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

//########################
//DingTalk类
//########################
type DingTalk struct {
	RobotUrl      string
	MessageStruct MessageStruct
}
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

//########################
//方式：MessageRenders
//########################
// MessageRenders 把模板渲染成消息体（返回 string 类型）
func (this DingTalk) MessageRenders() string {
	//创建一个字节缓冲器
	var buf bytes.Buffer
	//以 MessageTemplate 为蓝本，新建一个模板，起一个名字叫tmplNewName
	tmpl, err := template.New("tmplNewName").Parse(MessageTemplate)
	if err != nil {
		fmt.Println(err)
	}
	//实例化属性
	tmplData := this.MessageStruct
	//渲染tmpl模板，渲染的过程中把tmplData的值填充到模板之内，最后把渲染后的结果存入到buf这个字节缓冲器中
	if err = tmpl.Execute(&buf, tmplData); err != nil {
		panic(err)
	}
	//格式化字节缓冲器中的内容为String串，然后返回给调用方
	return buf.String()
}

//########################
//方法：MessagesJoins
//########################
// MessagesJoins 拼装成完整的报警消息（返回 map[string]interface{} 类型）
func (this DingTalk) MessagesJoins() map[string]interface{} {
	var MessageOk = map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": "Alertmanager已触发",
			"text":  this.MessageRenders(),
		},
	}
	return MessageOk
}

//########################
//方法：SendMessage
//########################
// SendMessage 向钉钉机器人发送信息
func (this DingTalk) SendMessage() ResponseData {
	url := this.RobotUrl
	data := this.MessagesJoins()
	contentType := "application/json;charset=utf-8"
	//结构体转换为json
	e, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	//发送post请求
	client := &http.Client{}
	response, err := client.Post(url, contentType, bytes.NewBuffer(e)) //post请求时，数据必须是byte
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println(err2)
	}

	req := ResponseData{}
	json.Unmarshal(body, &req)

	fmt.Println(string(body)) //打印的是人类可读的信息
	//return string(body)
	return req
}

//创建一个结构体用于接收dingding返回的参数
type ResponseData struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//实例化DingTalk对象
func NewDingTalk(url string, messageData MessageStruct) DingTalk {
	p := DingTalk{}
	p.RobotUrl = url
	p.MessageStruct = messageData
	return p
}
