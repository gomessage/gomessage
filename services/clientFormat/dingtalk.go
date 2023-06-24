package clientFormat

// PackDingtalkMessage 钉钉消息渲染（这是最终的需要发送出去的消息）包含一些符合对应客户端的工具字段
func PackDingtalkMessage(keyword string, message string) interface{} {
	type MarkdownMessage struct {
		Msgtype  string `json:"msgtype"`
		Markdown struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		} `json:"markdown"`
	}
	m := MarkdownMessage{}
	m.Msgtype = "markdown"
	m.Markdown.Title = "GoMessage：" + keyword
	m.Markdown.Text = message
	return m
}
