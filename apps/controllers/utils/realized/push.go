package realized

import (
	"encoding/json"
	"fmt"
	"gomessage/apps/controllers/clients"
	"gomessage/apps/controllers/send"
	"gomessage/apps/controllers/utils/base"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type GeneralPush struct {
	base.Push
}

func (d *GeneralPush) PushData(url string, data any) {
	fmt.Println("普通的post推送方法...")
	send.Push(d, url)
}

type WechatPush struct {
	base.Push
	CorpId      string
	AgentId     string
	AgentSecret string
	Touser      string
}

func (w *WechatPush) PushData(url string, data any) {
	byt, err := json.Marshal(data)

	msg := clients.PushMessageData{}
	msg.MsgType = "markdown"
	msg.Touser = w.Touser
	msg.AgentId, _ = strconv.Atoi(w.AgentId)
	msg.Markdown.Content = string(byt)

	MyByte, _ := json.Marshal(&msg)
	url = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + w.getAccessToken().AccessToken
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
	fmt.Println(string(body))
}

// 向微信发送请求获取access_token
func (w *WechatPush) getAccessToken() clients.GetAccessTokenReturn {
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
	r := clients.GetAccessTokenReturn{}
	json.Unmarshal(result, &r)
	return r
}
