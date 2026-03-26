package v2

import (
	"encoding/json"
	"fmt"
	"gomessage/pkg/services/core/v1"
	"gomessage/pkg/services/format"
	"gomessage/pkg/utils/log/loggers"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type GeneralPush struct {
	Push
}

func (d *GeneralPush) PushData(url string, data any) {
	fmt.Println("普通的post推送方法...")
	v1.Push(data, url)
}

type WechatPush struct {
	Push
	CorpId      string
	AgentId     string
	AgentSecret string
	Touser      string
}

func (w *WechatPush) PushData(url string, data any) {
	byt, err := json.Marshal(data)
	if err != nil {
		loggers.DefaultLogger.Errorln("微信应用消息序列化失败：", err)
		return
	}
	url = ""

	//要推送的数据
	msg := format.PushMessageData{}
	msg.MsgType = "markdown"
	msg.Touser = w.Touser
	msg.AgentId, err = strconv.Atoi(w.AgentId)
	if err != nil {
		loggers.DefaultLogger.Errorln("微信应用 agent_id 解析失败：", err)
		return
	}
	msg.Markdown.Content = string(byt)

	MyByte, err := json.Marshal(&msg)
	if err != nil {
		loggers.DefaultLogger.Errorln("微信应用请求体序列化失败：", err)
		return
	}
	accessToken := w.getAccessToken().AccessToken
	if accessToken == "" {
		loggers.DefaultLogger.Errorln("微信应用 access_token 为空")
		return
	}
	url = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + accessToken
	resp, err := http.Post(url, "application/json", strings.NewReader(string(MyByte)))
	if err != nil {
		loggers.DefaultLogger.Errorln("微信应用推送失败：", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		loggers.DefaultLogger.Errorln("微信应用响应读取失败：", err)
		_ = resp.Body.Close()
		return
	}
	if err = resp.Body.Close(); err != nil {
		loggers.DefaultLogger.Errorln("微信应用响应关闭失败：", err)
	}
	fmt.Println(string(body))
}

// 向微信发送请求获取access_token
func (w *WechatPush) getAccessToken() format.GetAccessTokenReturn {
	corpId := w.CorpId
	agentSecret := w.AgentSecret

	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpId + "&corpsecret=" + agentSecret
	resp, err := http.Get(url)
	if err != nil {
		loggers.DefaultLogger.Errorln("微信应用 token 请求失败：", err)
		return format.GetAccessTokenReturn{}
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		loggers.DefaultLogger.Errorln("微信应用 token 响应读取失败：", err)
		_ = resp.Body.Close()
		return format.GetAccessTokenReturn{}
	}
	if err = resp.Body.Close(); err != nil {
		loggers.DefaultLogger.Errorln("微信应用 token 响应关闭失败：", err)
	}
	r := format.GetAccessTokenReturn{}
	if err := json.Unmarshal(result, &r); err != nil {
		return format.GetAccessTokenReturn{}
	}
	return r
}
