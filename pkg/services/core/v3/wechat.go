package v3

import (
	"encoding/json"
	"fmt"
	"gomessage/pkg/models"
	"gomessage/pkg/services/core/v1"
	"gomessage/pkg/services/format"
	"gomessage/pkg/utils"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type ClientActionWechatApplication struct {
	CorpId      string
	AgentId     string
	AgentSecret string
	Touser      string
}

func (c *ClientActionWechatApplication) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	var msgList []any
	//是否聚合
	if isMerge {
		//把多个消息拼接成一个长字符串
		msg := v1.MessageJoint(contentList, utils.VarWechatApplication)

		//把普通的内容体渲染成符合微信应用号的消息体
		message := format.PushMessageData{}
		message.MsgType = "markdown"
		message.Touser = c.Touser
		message.AgentId, _ = strconv.Atoi(c.AgentId)
		message.Markdown.Content = msg

		msgList = append(msgList, message)
	} else {
		for _, msg := range contentList {

			//把普通的内容体渲染成符合微信应用号的消息体
			message := format.PushMessageData{}
			message.MsgType = "markdown"
			message.Touser = c.Touser
			message.AgentId, _ = strconv.Atoi(c.AgentId)
			message.Markdown.Content = msg

			msgList = append(msgList, message)
		}
	}
	return msgList
}

func (c *ClientActionWechatApplication) PushMessages(messages []any) {
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
func (c *ClientActionWechatApplication) getAccessToken() format.GetAccessTokenReturn {
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
	r := format.GetAccessTokenReturn{}
	json.Unmarshal(result, &r)
	return r
}
