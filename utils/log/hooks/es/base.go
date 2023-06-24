package es

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type Record struct {
	Timestamp string        `json:"@timestamp"`        //时间戳
	Host      string        `json:"Host,omitempty"`    //主机（omitempty：如果不存在则忽略，连空字符串都不显示）
	File      string        `json:"File,omitempty"`    //文件名（omitempty：如果不存在则忽略，连空字符串都不显示）
	Func      string        `json:"Func,omitempty"`    //函数名（omitempty：如果不存在则忽略，连空字符串都不显示）
	Message   string        `json:"Message,omitempty"` //日志内容（omitempty：如果不存在则忽略，连空字符串都不显示）
	Level     string        `json:"Level,omitempty"`   //日志级别（omitempty：如果不存在则忽略，连空字符串都不显示）
	Data      logrus.Fields `json:"data,omitempty"`
}

func CreateRecord(entry *logrus.Entry) *Record {
	level := entry.Level.String()
	if e, ok := entry.Data[logrus.ErrorKey]; ok && e != nil {
		if err, ok := e.(error); ok {
			entry.Data[logrus.ErrorKey] = err.Error()
		}
	}
	var file string
	var function string
	if entry.HasCaller() {
		file = entry.Caller.File
		function = entry.Caller.Function
	}
	//获取本机host
	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}
	return &Record{
		Host:      host,
		Timestamp: entry.Time.UTC().Format(time.RFC3339Nano),
		File:      file,
		Func:      function,
		Message:   entry.Message,
		Data:      entry.Data,
		Level:     strings.ToUpper(level), //转大写
	}
}

func SyncIndex(indexName string) {
	//设置上下文超时时间
	ctx := context.Background()

	//判断索引是否存在，如果不存在则创建
	exists, err := elasticsearchClient.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	} else {
		if !exists {
			_, err := elasticsearchClient.CreateIndex(indexName).BodyString(Mapping2).Do(ctx)
			if err != nil {
				panic(err)
			}
		}
	}
}

func JoinIndexName(indexName string) string {
	return fmt.Sprintf("%s-%s", indexName, time.Now().Format("2006.01.02"))
}
