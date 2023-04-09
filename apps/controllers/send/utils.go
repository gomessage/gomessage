package send

import (
	"fmt"
	"strings"
)

// MessageJoint 按照不同的客户端类型，选择不同的拼接和聚合形式
func MessageJoint(messageList []string, thisType string) string {
	var msg string
	if thisType == "dingtalk" {
		msg = strings.Join(messageList, "\n * * * \n")

	} else if thisType == "wechat" {
		msg = strings.Join(messageList, "\n · \n")

	} else if thisType == "feishu" {
		msg = strings.Join(messageList, "\n -------------- \n")

	} else {
		fmt.Println("客户端类型错误")
	}
	return msg
}
