package WebMessageSend

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//获取access_token时返回值的结构体
type GetAccessTokenReturn struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

//要推送的消息体的结构
type PushMessageData struct {
	Touser   string `json:"touser"`
	Msgtype  string `json:"msgtype"`
	Agentid  int    `json:"agentid"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

//推送message信息时返回值的结构体
//type PushMessageReturn struct {
//	Errcode int    `json:"errcode"`
//	Errmsg  string `json:"errmsg"`
//	Msgid   string `json:"msgid"`
//}

//########################
//WeChat类
//########################
type WeChat struct {
	access_token_url string
	push_message_url string
	Corpid           string
	Agentid          string
	Agentsecret      string
	Touser           string
	Content          string
}

//想微信发送请求获取access_token
func (this WeChat) getAccessToken() GetAccessTokenReturn {
	tmpCorpId := this.Corpid
	tmpAgentsecret := this.Agentsecret
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + tmpCorpId + "&corpsecret=" + tmpAgentsecret
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	result, err := ioutil.ReadAll(resp.Body)
	r := GetAccessTokenReturn{}
	json.Unmarshal(result, &r)
	return r
}

//实际推送信息的方法
func (this WeChat) PushMessage() {
	//要发送出去的数据体
	msg := PushMessageData{}
	msg.Msgtype = "markdown"
	msg.Touser = this.Touser
	msg.Agentid, _ = strconv.Atoi(this.Agentid)
	msg.Markdown.Content = this.Content

	MyByte, _ := json.Marshal(&msg)
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + this.getAccessToken().AccessToken
	resp, err := http.Post(url, "application/json", strings.NewReader(string(MyByte)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//r := PushMessageReturn{}
	//json.Unmarshal(result, &r)
	//return r
	fmt.Println(string(body))
}
