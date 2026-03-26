package format

import (
	"encoding/json"
	"fmt"
	"gomessage/pkg/utils/log/loggers"
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
		loggers.DefaultLogger.Errorln("企业微信 token 请求失败：", err)
		return GetAccessTokenReturn{}
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		loggers.DefaultLogger.Errorln("企业微信 token 响应读取失败：", err)
		_ = resp.Body.Close()
		return GetAccessTokenReturn{}
	}
	if err = resp.Body.Close(); err != nil {
		loggers.DefaultLogger.Errorln("企业微信 token 响应关闭失败：", err)
	}
	r := GetAccessTokenReturn{}
	if err = json.Unmarshal(result, &r); err != nil {
		loggers.DefaultLogger.Errorln("企业微信 token 响应解析失败：", err)
		return GetAccessTokenReturn{}
	}
	return r
}

// PushMessage 实际推送信息的方法
func (w *WeChat) PushMessage() {
	//要发送出去的数据体
	msg := PushMessageData{}
	msg.MsgType = "markdown"
	msg.Touser = w.Touser
	agentID, err := strconv.Atoi(w.AgentId)
	if err != nil {
		loggers.DefaultLogger.Errorln("企业微信 agent_id 解析失败：", err)
		return
	}
	msg.AgentId = agentID
	msg.Markdown.Content = w.Content

	MyByte, err := json.Marshal(&msg)
	if err != nil {
		loggers.DefaultLogger.Errorln("企业微信消息序列化失败：", err)
		return
	}
	token := w.getAccessToken().AccessToken
	if token == "" {
		loggers.DefaultLogger.Errorln("企业微信 access_token 为空")
		return
	}
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + token
	resp, err := http.Post(url, "application/json", strings.NewReader(string(MyByte)))
	if err != nil {
		loggers.DefaultLogger.Errorln("企业微信消息推送失败：", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		loggers.DefaultLogger.Errorln("企业微信响应读取失败：", err)
		_ = resp.Body.Close()
		return
	}
	if err = resp.Body.Close(); err != nil {
		loggers.DefaultLogger.Errorln("企业微信响应关闭失败：", err)
	}
	//r := PushMessageReturn{}
	//json.Unmarshal(result, &r)
	//return r
	fmt.Println(string(body))
}
