package web2

import (
	"GoMessage/models"
	"bytes"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/tidwall/gjson"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type ApiControllers struct {
	beego.Controller
}

// @router /message [get]
func (this *ApiControllers) Get() {
	//给出返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = "ok"
	this.ServeJSON()
}

// @router /message [post]
func (this *ApiControllers) Post() {
	//进行CacheData数据绑定
	CacheData.RequestBody = this.Ctx.Input.RequestBody
	json.Unmarshal(CacheData.RequestBody, &CacheData.MessageData)
	CacheData.UpdateTime = time.Now()

	//拿到推送过来的[]byte消息（OK）
	messageByteData := CacheData.RequestBody

	//从数据库中拿到"模板变量"映射关系 （数据库方法）
	//从数据库中拿到"模板内容" （数据库方法）
	//从数据库中拿到钉钉客户端的配置 （数据库方法）
	conf := GetConfig()

	//获得自定义变量与数据的映射 AnalysisData()
	analysisData := AnalysisData(conf.mapConfig, messageByteData)

	//通过自定义变量+模板内容渲染出最终的消息体 MessageRenders()
	message := MessageRenders(conf.templateConfig, analysisData)
	fmt.Println("渲染后的信息：" + message)

	//把最终的消息体通过钉钉发送出去
	result := SendMessage(conf.dingtalkConfig, message)
	fmt.Println(result)

	//给出返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = "ok"
	this.ServeJSON()
}

func DingtalkRandomUrl(dingding []models.DingtalkClient) models.DingtalkClient {
	dingtalkClient := dingding
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(dingtalkClient)
	return dingtalkClient[n]
}

//========================
//消息体的最终格式化
//========================
func formatMessage(keyword string, message string) interface{} {
	type MarkdownMessage struct {
		Msgtype  string `json:"msgtype"`
		Markdown struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		} `json:"markdown"`
	}
	m := MarkdownMessage{}
	m.Msgtype = "markdown"
	m.Markdown.Title = "GoMessage：" + keyword
	m.Markdown.Text = message
	return m
}

//========================
//发送数据的实际方法
//========================
func SendMessage(dingding []models.DingtalkClient, message string) interface{} {
	dingtalkCli := DingtalkRandomUrl(dingding)
	url := dingtalkCli.RobotUrl
	keyword := dingtalkCli.RobotKeyword
	data := formatMessage(keyword, message)

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
	type DingtalkResponse struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}
	req := DingtalkResponse{}
	json.Unmarshal(body, &req)

	fmt.Println(string(body)) //打印的是人类可读的信息
	//return string(body)
	return req
}

type Config struct {
	mapConfig      []map[string]string
	templateConfig string
	dingtalkConfig []models.DingtalkClient
}

//========================
//获取到数据库中存储的用户配置
//========================
func GetConfig() Config {
	c := Config{}

	//获取Map相关的配置
	var tmpList []map[string]string
	for _, value := range models.QueryAllMap() {
		tmpMap := make(map[string]string)
		k := value.Key
		v := value.Value
		tmpMap[k] = v
		tmpList = append(tmpList, tmpMap)
	}
	c.mapConfig = tmpList

	//获取Template相关的配置
	c.templateConfig = models.GetOneTemplate("default").MessageTemplate

	//获取dingding相关的配置
	c.dingtalkConfig = models.QueryAllDingtalkClient()

	return c
}

//========================
//模板变量与实际的RequestData对应起来，建立映射关系
//========================
func AnalysisData(keyValueList []map[string]string, messageData []byte) map[string]string {
	var dataMapping map[string]string
	dataMapping = make(map[string]string)
	var tmpKey string
	var tmpValue gjson.Result
	for _, dict := range keyValueList {
		for k, v := range dict {
			tmpKey = v
			tmpValue = gjson.GetBytes(messageData, tmpKey)
			//fmt.Println(k, tmpValue)
			dataMapping[k] = tmpValue.String()
		}
	}
	return dataMapping
}

//========================
//接收用户的模板信息，然后加上映射好的数据模型，直接返回一个完整的信息体出来
//========================
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
