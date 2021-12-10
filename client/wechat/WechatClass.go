package wechat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//获取access_token时返回值的结构体
type GetAccessTokenReturn struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

//推送message信息时返回值的结构体
type PushMessageReturn struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Msgid   string `json:"msgid"`
}

//要推送的消息体的结构
type PushMessageData struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe int `json:"safe"`
}

//########################
//WeChat类
//########################
type WeChat struct {
	corpid           string
	corpsecret       string
	access_token_url string
	push_message_url string
}

//func (this WeChat) init() {
//	this.corpid = "ww8290f080960950c4"
//	this.corpsecret = "5mtW-dbOiTXPYPD3VidAZq1EfK3h5nGIp8B0zWuA4z0"
//	this.access_token_url = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + this.corpid + "&corpsecret=" + this.corpsecret
//	//this.push_message_url = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + this.getAccessToken().AccessToken
//}

//想微信发送请求获取access_token
func (this WeChat) getAccessToken() GetAccessTokenReturn {

	corpid := "ww8290f080960950c4"
	corpsecret := "5mtW-dbOiTXPYPD3VidAZq1EfK3h5nGIp8B0zWuA4z0"
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpid + "&corpsecret=" + corpsecret

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)

	r := GetAccessTokenReturn{}
	json.Unmarshal(result, &r)
	return r
}

func (this WeChat) PushMessage(content string) PushMessageReturn {

	//初始化自身
	//this.init()
	//this.getAccessToken()

	msg := PushMessageData{}
	msg.Touser = "000001|000002"
	msg.Msgtype = "text"
	msg.Agentid = 1000010
	msg.Text.Content = content
	msg.Safe = 0

	MyByte, _ := json.Marshal(&msg)
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + this.getAccessToken().AccessToken
	resp, err := http.Post(url, "application/json", strings.NewReader(string(MyByte)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	r := PushMessageReturn{}
	json.Unmarshal(result, &r)

	return r
}
