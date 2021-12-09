package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

const MessageTemplate = `
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

//DingTalk类
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

	//fmt.Println(tmpl.Name())  //打印模板的名字
	//fmt.Println(buf.String()) //打印字节缓冲器中的内容

	//格式化字节缓冲器中的内容为String串，然后返回给调用方
	return buf.String()
}

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

// SendMessage 向钉钉机器人发送信息
func (this DingTalk) SendMessage() {
	url := this.RobotUrl
	data := this.MessagesJoins()
	contentType := "application/json;charset=utf-8"

	fmt.Printf("========%v\n", data)
	fmt.Printf("========%v\n", url)

	//结构体转换为json
	e, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

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
	fmt.Println(string(body)) //打印的是人类可读的信息
	fmt.Println(body)         //打印的是二进制数据
}

//实例化DingTalk对象
func NewDingTalk(url string, template string, message MessageStruct) DingTalk {
	p := DingTalk{}
	p.RobotUrl = url
	p.MessageStruct = message
	fmt.Println(template)
	return p
}
