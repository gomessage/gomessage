package clientFormat

import "fmt"

// PackWechatRobotMessage 企业微信机器人消息渲染（这是最终的需要发送出去的消息）包含一些符合对应客户端的工具字段
func PackWechatRobotMessage(keyword string, message string) interface{} {
	fmt.Println(keyword)

	type MarkdownMessage struct {
		Msgtype  string `json:"msgtype"`
		Markdown struct {
			Content string `json:"content"`
		} `json:"markdown"`
	}

	m := MarkdownMessage{}
	m.Msgtype = "markdown"
	m.Markdown.Content = message

	return m
}
