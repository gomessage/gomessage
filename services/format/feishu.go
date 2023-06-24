package format

import (
	"gomessage/models"
)

// PackFeishuMessage 飞书消息渲染（这是最终的需要发送出去的消息）包含一些符合对应客户端的工具字段
func PackFeishuMessage(userConfigInfo *models.Client, message string) interface{} {
	//fmt.Println(keyword)

	//普通的文本格式消息
	//type Text struct {
	//	Tag  string `json:"tag"`
	//	Text struct {
	//		Lines   int    `json:"lines"`
	//		Tag     string `json:"tag"`
	//		Content string `json:"Content"`
	//	} `json:"text"`
	//}

	//markdown格式的消息
	type Markdown struct {
		Lines   int    `json:"lines"`
		Tag     string `json:"tag"`
		Content string `json:"Content"`
	}

	//最终的结构体
	type TextMessage struct {
		MsgType string `json:"msg_type"`
		Card    struct {
			Header struct {
				Title struct {
					Tag     string `json:"tag"`
					Content string `json:"content"`
				} `json:"title"`
				Template string `json:"template"`
			} `json:"header"`
			Elements []Markdown `json:"elements"`
		} `json:"card"`
	}

	//实例化消息体
	var mList []Markdown
	m := Markdown{}
	m.Lines = 200
	m.Tag = "markdown"
	m.Content = message
	mList = append(mList, m)

	//拼装数据体
	t := TextMessage{}
	t.MsgType = "interactive"
	t.Card.Header.Template = userConfigInfo.ExtendFeishu.TitleColor
	t.Card.Header.Title.Content = userConfigInfo.ExtendFeishu.RobotKeyword
	t.Card.Header.Title.Tag = "plain_text"
	t.Card.Elements = mList

	return t
}
