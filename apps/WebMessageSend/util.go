package WebMessageSend

import (
	"GoMessage/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"html/template"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//========================
//结构体：存放用户配置的实例化
//========================
type Config struct {
	ConfigMap       []map[string]string
	MessageTemplate string
	MessageMerge    bool
	ActiveClient    []models.Client
}

//========================
//获取到用户在图形界面上设置的各种参数和变量
//========================
func GetUserConfig() Config {
	c := Config{}

	//获取ConfigMap相关的配置（用户的变量映射）
	var tmpList []map[string]string
	for _, value := range models.QueryAllMap() {
		tmpMap := make(map[string]string)
		k := value.Key
		v := value.Value
		tmpMap[k] = v
		tmpList = append(tmpList, tmpMap)
	}
	c.ConfigMap = tmpList

	//获取Template相关的配置（这里拿到的是用户的模板骨架）
	templateObject := models.GetOneTemplate("default")
	c.MessageTemplate = templateObject.MessageTemplate

	//（这里拿到的是消息是否聚合）
	c.MessageMerge = templateObject.MessageMerge

	//获取客户端相关（只获取激活的客户端）
	c.ActiveClient = models.GetClentActive()
	return c
}

//========================
//随机获取一个机器人地址（通用方法可以同时被钉钉和飞书使用）
//========================
func RobotRandomUrl(dList []string) string {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(dList)
	return dList[n]
}

//========================
//判断客户端
//========================
//func VerdictClient(activeClient *models.Client, message string) (interface{}, string) {
//	//判断激活客户端的类型
//	if activeClient.Type == "dingtalk" {
//		keyword := activeClient.ClientDingtalk.RobotKeyword             //获得机器人放行关键字
//		url := RobotRandomUrl(activeClient.ClientDingtalk.RobotUrlList) //随机获得一个机器人地址
//		data := MessageRendersDingtalk(keyword, message)
//		return data, url
//
//	} else if activeClient.Type == "wechat" {
//		fmt.Println("暂时先不发送消息")
//
//	} else if activeClient.Type == "feishu" {
//		url := RobotRandomUrl(activeClient.ClientFeishu.RobotUrlList)
//		data := MessageRendersFeishu(activeClient, message)
//		return data, url
//
//	} else {
//		fmt.Println("客户端类型错误，不会发送消息")
//	}
//	return "ok", "ok"
//}

//========================
//真正发送消息的方法，不错任何形式的数据处理，仅仅只是单纯的发送
//========================
func SendMessage(data interface{}, url string) {
	contentType := "application/json;charset=utf-8"

	//结构体转换为json
	e, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	//发送post请求
	client := &http.Client{}
	//response, err := client.Post(url, contentType, bytes.NewBuffer(e)) //post请求时，数据必须是byte
	response, err := client.Post(url, contentType, strings.NewReader(string(e))) //post请求时，数据必须是byte
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(string(body)) //打印的是人类可读的信息
}

//========================
//企业微信消息渲染（这是最终的需要发送出去的消息）包含一些符合对应客户端的工具字段
//========================
func MessageRendersWechat(keyword string, message string) interface{} {
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
//飞书消息渲染（这是最终的需要发送出去的消息）包含一些符合对应客户端的工具字段
//========================
func MessageRendersFeishu(userConfigInfo *models.Client, message string) interface{} {
	//fmt.Println(keyword)

	//普通的文本格式消息
	//type Text struct {
	//	Tag  string `json:"tag"`
	//	Text struct {
	//		Lines   int    `json:"lines"`
	//		Tag     string `json:"tag"`
	//		Content string `json:"Content"`
	//	} `json:"text"`
	//}

	//markdown格式的消息
	type Markdown struct {
		Lines   int    `json:"lines"`
		Tag     string `json:"tag"`
		Content string `json:"Content"`
	}

	//最终的结构体
	type TextMessage struct {
		MsgType string `json:"msg_type"`
		Card    struct {
			Header struct {
				Title struct {
					Tag     string `json:"tag"`
					Content string `json:"content"`
				} `json:"title"`
				Template string `json:"template"`
			} `json:"header"`
			Elements []Markdown `json:"elements"`
		} `json:"card"`
	}

	//实例化消息体
	var mList []Markdown
	m := Markdown{}
	m.Lines = 200
	m.Tag = "markdown"
	m.Content = message
	mList = append(mList, m)

	//拼装数据体
	t := TextMessage{}
	t.MsgType = "interactive"
	t.Card.Header.Template = userConfigInfo.ClientFeishu.TitleColor
	t.Card.Header.Title.Content = userConfigInfo.ClientFeishu.RobotKeyword
	t.Card.Header.Title.Tag = "plain_text"
	t.Card.Elements = mList

	return t
}

//========================
//钉钉消息渲染（这是最终的需要发送出去的消息）包含一些符合对应客户端的工具字段
//========================
func MessageRendersDingtalk(keyword string, message string) interface{} {
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
//模板渲染（这仅仅只是把模板中内嵌的变量渲染成消息实体）
//========================
func TemplateRenders(thisTemplate string, dataList []map[string]string) []string {

	var templateList []string

	for _, data := range dataList {

		//创建一个字节缓冲器
		var buf bytes.Buffer

		//以 K8sMessageTemplate 为蓝本，新建一个模板，新起一个名字叫tmplNewName
		tmpl, err := template.New("tmplNewName").Parse(thisTemplate)
		if err != nil {
			fmt.Println(err)
		}
		//实例化属性
		//tmplData := data
		//fmt.Println(data)

		//渲染tmpl模板，渲染的过程中把tmplData的值填充到模板之内，最后把渲染后的结果存入到buf这个字节缓冲器中
		if err = tmpl.Execute(&buf, data); err != nil {
			panic(err)
		}

		//格式化字节缓冲器中的内容为String串，然后返回给调用方
		//return buf.String()
		templateList = append(templateList, buf.String())
	}
	return templateList
}

//========================
//模板变量与实际的RequestData对应起来，建立映射关系
//========================
func AnalysisData(keyValueList []map[string]string, messageData []byte) []map[string]string {

	//这一段代码，只是从原始数据中拿到【用户需求字段】中长度最长的那一个度量单位，不对原始数据做任何处理和拆分
	//同时这段代码，只关注用户关注的变量，其它多余的变量都会被舍弃
	//再再同时，这段代码只关注变量对应的value，不关注key的处理
	lenNum := 0
	for _, dict := range keyValueList {
		for _, v := range dict {
			tmpKey := v                                     //此时v里面存放的内容是这种形式的：alert.#.label.name
			tmpValue := gjson.GetBytes(messageData, tmpKey) //gjson支持 aaa.#.bbb这种形式的取值，取到的内容是：[bbb1,bbb2,bbb3,bbb4,bbb5]这种形式的
			if len(tmpValue.Array()) >= lenNum {            //拿到当前这个子元素中，最长数据的长度，这个数字用来做下文遍历使用
				lenNum = len(tmpValue.Array()) //让for循环外面声明的那个lenNum的长度，达到当前实际取值中的最大值（而且只关注用户关注的值），用户不关注的值我们也不去操心
			}
		}
	}

	//下面这个代码，会根绝lenNum的长度遍历多次，每一次都拼装好一个完整的dict，然后追加进dataList中。
	var dataList []map[string]string
	for i := 0; i < lenNum; i++ {
		oneDataMapping := make(map[string]string)
		for _, dict := range keyValueList {
			for k, v := range dict {
				tmpKey := strings.ReplaceAll(v, "#", strconv.Itoa(i))
				//fmt.Println(tmpKey)
				tmpValue := gjson.GetBytes(messageData, tmpKey) //此时的：tmpKey是：【alert.0.label.name】 这种形式，取到的是一个具体值
				oneDataMapping[k] = tmpValue.String()
			}
		}
		//fmt.Println(oneDataMapping)
		dataList = append(dataList, oneDataMapping)
	}

	//返回一个list，里面包含了若干个dict类型
	return dataList
}

//按照不同的客户端类型，选择不同的拼接和聚合形式
func MessageJoint(messageList []string, thisType string) string {
	var msg string
	if thisType == "dingtalk" {
		msg = strings.Join(messageList, "\n * * * \n")

	} else if thisType == "wechat" {
		msg = strings.Join(messageList, "\n · \n")

	} else if thisType == "feishu" {
		msg = strings.Join(messageList, "\n -------------- \n")

	} else {
		fmt.Println("客户端类型错误")
	}
	return msg
}
