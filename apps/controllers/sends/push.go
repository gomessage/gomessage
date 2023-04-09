package sends

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gomessage/utils/log/loggers"
	"io"
	"net/http"
	"strings"
	"time"
)

// SendMessage 真正发送消息的方法，不做任何形式的数据处理，仅仅只是单纯的发送
func SendMessage(data interface{}, url string) {
	contentType := "application/json;charset=utf-8"

	//结构体转换为json
	e, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	//发送post请求
	client := &http.Client{}
	//response, err := client.Post(url, contentType, bytes.NewBuffer(e)) //post请求时，数据必须是byte
	response, err := client.Post(url, contentType, strings.NewReader(string(e)))
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println(err2)
	}

	loggers.PushLogger.WithFields(logrus.Fields{
		"content_type":    contentType,
		"url":             url,
		"request_body":    string(e),
		"response_body":   string(body),
		"response_status": response.Status,
		"time_now":        time.Now().Format("2006-01-02_15:04:05.000000"),
	}).Info("推送数据成功")

	fmt.Println(string(body)) //打印的是人类可读的信息
}
