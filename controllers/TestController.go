package controllers

import (
	web2 "GoMessage/controllers/web2"
	"bytes"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/tidwall/gjson"
	"text/template"
)

type TestController struct {
	beego.Controller
}

// @router / [get]
func (this *TestController) Get() {
	fmt.Println(web2.CacheData)

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = "hello"
	this.ServeJSON()
}

// @router / [post]
func (this *TestController) Post() {
	//web2.CacheData.RequestBody = this.Ctx.Input.RequestBody
	messageData := web2.CacheData.RequestBody
	//fmt.Println(messageData)

	type requestData struct {
		KeyValueList []map[string]string `json:"key_value_list"`
	}
	Config := requestData{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &Config)

	if messageData != nil {
		tmpaaa := analysisData(Config.KeyValueList, messageData)
		tmpbbb := MessageRenders(TestMessageTemplate, tmpaaa)
		fmt.Println(tmpbbb)
	} else {
		fmt.Println("messageData中没有数据...")
	}

	//var tmpKey string
	//var tmpValue gjson.Result
	//for _, dict := range Config.KeyValueList {
	//	for k, v := range dict {
	//		//fmt.Println(k, v)
	//		tmpKey = v
	//		tmpValue = gjson.GetBytes(messageData, tmpKey)
	//		fmt.Println(k, v)
	//		fmt.Println(k, tmpValue)
	//	}
	//}

	//fmt.Println(Config)
	//aaa := gjson.GetBytes(messageData, "alerts.0.labels.alertname")
	//fmt.Println("============", aaa)

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = Config
	this.ServeJSON()
}

//解析数据，把用户创建的配置变量，与实际的RequestData对应起来，建立映射关系
func analysisData(keyValueList []map[string]string, messageData []byte) map[string]string {
	var dataMapping map[string]string
	dataMapping = make(map[string]string)
	var tmpKey string
	var tmpValue gjson.Result
	for _, dict := range keyValueList {
		for k, v := range dict {
			tmpKey = v
			tmpValue = gjson.GetBytes(messageData, tmpKey)
			fmt.Println(k, tmpValue)
			dataMapping[k] = tmpValue.String()
		}
	}
	return dataMapping
}

//接收用户的模板信息，然后加上映射好的数据模型，直接返回一个完整的信息体出来
func MessageRenders(thisTemplate string, data map[string]string) string {
	//创建一个字节缓冲器
	var buf bytes.Buffer
	//以 K8sMessageTemplate 为蓝本，新建一个模板，新起一个名字叫tmplNewName
	tmpl, err := template.New("tmplNewName").Parse(thisTemplate)
	if err != nil {
		fmt.Println(err)
	}
	//实例化属性
	tmplData := data
	//渲染tmpl模板，渲染的过程中把tmplData的值填充到模板之内，最后把渲染后的结果存入到buf这个字节缓冲器中
	if err = tmpl.Execute(&buf, tmplData); err != nil {
		panic(err)
	}
	//格式化字节缓冲器中的内容为String串，然后返回给调用方
	return buf.String()
}

//########################
//template：用来拼装出来要转发出去的报警信息
//########################
const TestMessageTemplate = `

**告警规则**：{{ .alertName }}

**告警级别**：{{ .aaa }}

**主机名称**：{{ .bbb }} 

**主机地址**：{{ .ccc }}

`
