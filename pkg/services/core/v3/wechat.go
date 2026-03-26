package v3

import (
	"encoding/json"
	"fmt"
	"gomessage/pkg/models"
	v1 "gomessage/pkg/services/core/v1"
	"gomessage/pkg/services/format"
	"gomessage/pkg/utils"
	"gomessage/pkg/utils/log/loggers"
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

func (c *ClientActionWechatApplication) PushMessages(messages []any) (successCount int, failureCount int) {
	token := c.getAccessToken()
	if token.AccessToken == "" {
		loggers.DefaultLogger.Error("企业微信 access_token 获取失败")
		failureCount = len(messages)
		return
	}
	for _, msg2 := range messages {
		MyByte, err := json.Marshal(msg2)
		if err != nil {
			loggers.DefaultLogger.Errorln("企业微信消息序列化失败：", err)
			failureCount++
			continue
		}
		url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + token.AccessToken
		resp, err := http.Post(url, "application/json", strings.NewReader(string(MyByte)))
		if err != nil {
			loggers.DefaultLogger.Errorln("企业微信消息推送失败：", err)
			failureCount++
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			loggers.DefaultLogger.Errorln("企业微信响应读取失败：", err)
			_ = resp.Body.Close()
			failureCount++
			continue
		}
		if err = resp.Body.Close(); err != nil {
			loggers.DefaultLogger.Errorln("企业微信响应关闭失败：", err)
		}
		fmt.Println(string(body))
		successCount++
	}
	return
}

// 向微信发送请求获取access_token
func (c *ClientActionWechatApplication) getAccessToken() format.GetAccessTokenReturn {
	corpId := c.CorpId
	agentSecret := c.AgentSecret

	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpId + "&corpsecret=" + agentSecret
	resp, err := http.Get(url)
	if err != nil {
		loggers.DefaultLogger.Errorln("企业微信 token 请求失败：", err)
		return format.GetAccessTokenReturn{}
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		loggers.DefaultLogger.Errorln("企业微信 token 响应读取失败：", err)
		_ = resp.Body.Close()
		return format.GetAccessTokenReturn{}
	}
	if err = resp.Body.Close(); err != nil {
		loggers.DefaultLogger.Errorln("企业微信 token 响应关闭失败：", err)
	}
	r := format.GetAccessTokenReturn{}
	if err = json.Unmarshal(result, &r); err != nil {
		loggers.DefaultLogger.Errorln("企业微信 token 响应解析失败：", err)
		return format.GetAccessTokenReturn{}
	}
	return r
}
