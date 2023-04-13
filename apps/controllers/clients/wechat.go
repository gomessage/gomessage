package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// GetAccessTokenReturn 获取access_token时返回值的结构体
type GetAccessTokenReturn struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// PushMessageData 要推送的消息体的结构
type PushMessageData struct {
	Touser   string `json:"touser"`
	MsgType  string `json:"msgtype"`
	AgentId  int    `json:"agentid"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

type WeChat struct {
	//access_token_url string
	//push_message_url string
	CorpId      string
	AgentId     string
	AgentSecret string
	Touser      string
	Content     string
}

// 向微信发送请求获取access_token
func (w *WeChat) getAccessToken() GetAccessTokenReturn {
	corpId := w.CorpId
	agentSecret := w.AgentSecret

	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpId + "&corpsecret=" + agentSecret
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

	result, err := io.ReadAll(resp.Body)
	r := GetAccessTokenReturn{}
	json.Unmarshal(result, &r)
	return r
}

// PushMessage 实际推送信息的方法
func (w *WeChat) PushMessage() {
	//要发送出去的数据体
	msg := PushMessageData{}
	msg.MsgType = "markdown"
	msg.Touser = w.Touser
	msg.AgentId, _ = strconv.Atoi(w.AgentId)
	msg.Markdown.Content = w.Content

	MyByte, _ := json.Marshal(&msg)
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + w.getAccessToken().AccessToken
	resp, err := http.Post(url, "application/json", strings.NewReader(string(MyByte)))
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	//r := PushMessageReturn{}
	//json.Unmarshal(result, &r)
	//return r
	fmt.Println(string(body))
}
