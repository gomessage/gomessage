package v1

import (
	"encoding/json"
	"fmt"
	"gomessage/pkg/utils/log/loggers"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Push 真正发送消息的方法，不做任何形式的数据处理，仅仅只是单纯的发送
func Push(data any, url string) error {
	contentType := "application/json;charset=utf-8"

	//结构体转换为json
	e, err := json.Marshal(data)
	if err != nil {
		loggers.DefaultLogger.Errorln("推送数据序列化失败：", err)
		return err
	}

	//发送post请求
	client := &http.Client{}
	//response, err := client.Post(url, contentType, bytes.NewBuffer(e))
	response, err := client.Post(url, contentType, strings.NewReader(string(e)))
	if err != nil {
		loggers.DefaultLogger.Errorln("推送请求失败：", err)
		return err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		loggers.DefaultLogger.Errorln("推送响应读取失败：", err)
		_ = response.Body.Close()
		return err
	}
	if err = response.Body.Close(); err != nil {
		loggers.DefaultLogger.Errorln("推送响应关闭失败：", err)
	}

	loggers.PushLogger.WithFields(logrus.Fields{
		"content_type":    contentType,
		"url":             url,
		"request_body":    string(e),
		"response_body":   string(body),
		"response_status": response.Status,
		"time_now":        time.Now().Format("2006-01-02_15:04:05.000000"),
	}).Info("推送数据成功")

	//打印人类可读的信息
	fmt.Println(string(body))
	return nil
}
