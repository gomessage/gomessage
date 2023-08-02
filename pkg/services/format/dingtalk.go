package format

func PackDingtalkMessage(keyword string, message string, atAll bool, mobiles []string) interface{} {
	type MarkdownMessage struct {
		Msgtype  string `json:"msgtype"`
		Markdown struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		} `json:"markdown"`
		At struct {
			AtMobiles []string `json:"atMobiles"`
			IsAtAll   bool     `json:"atAll"`
		} `json:"at"`
	}
	m := MarkdownMessage{}
	m.Msgtype = "markdown"
	m.Markdown.Title = keyword + "GoMessage"
	m.Markdown.Text = message

	if atAll { //判断是否要@所有人，如果为true，则人员名单设置为空
		m.At.IsAtAll = true
		m.At.AtMobiles = nil
	} else {
		m.At.IsAtAll = false
		m.At.AtMobiles = mobiles
	}

	return m
}
