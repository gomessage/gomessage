package clients

import (
	"encoding/json"
	"fmt"
	"gomessage/apps/controllers/clients"
	"gomessage/apps/controllers/send"
	"gomessage/apps/models"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type ClientActionWechat struct {
	CorpId      string
	AgentId     string
	AgentSecret string
	Touser      string
}

func (c *ClientActionWechat) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	//是否聚合
	if isMerge {
		//把多个消息拼接成一个长字符串
		msg := send.MessageJoint(contentList, "wechat")

		//把普通的内容体渲染成符合微信应用号的消息体
		message := clients.PushMessageData{}
		message.MsgType = "markdown"
		message.Touser = c.Touser
		message.AgentId, _ = strconv.Atoi(c.AgentId)
		message.Markdown.Content = msg

		msgList = append(msgList, message)
	} else {
		for _, msg := range contentList {

			//把普通的内容体渲染成符合微信应用号的消息体
			message := clients.PushMessageData{}
			message.MsgType = "markdown"
			message.Touser = c.Touser
			message.AgentId, _ = strconv.Atoi(c.AgentId)
			message.Markdown.Content = msg

			msgList = append(msgList, message)
		}
	}
	return msgList
}

func (c *ClientActionWechat) PushMessages(messages []any) {
	for _, msg2 := range messages {
		MyByte, _ := json.Marshal(msg2)
		url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + c.getAccessToken().AccessToken
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
}

// 向微信发送请求获取access_token
func (c *ClientActionWechat) getAccessToken() clients.GetAccessTokenReturn {
	corpId := c.CorpId
	agentSecret := c.AgentSecret

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
