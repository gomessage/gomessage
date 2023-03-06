package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"time"
)

var ElasticsearchClient *elastic.Client
var IndexName string
var MyHost string

func init() {
	var esUrl = "http://esxxxxxxxxxxxxxxxxx:9200"
	var esUserName = "xxxxxx"
	var esPassword = "xxxxxxxxxxx"
	var err error
	MyHost, err = os.Hostname()
	if err != nil {
		MyHost = "localhost"
	}

	ElasticsearchClient, err = elastic.NewClient(
		elastic.SetURL(esUrl),
		elastic.SetBasicAuth(esUserName, esPassword),
		elastic.SetGzip(true),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC：", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetSniff(false),
	)

	if err != nil {
		fmt.Println("ES连接失败：", err)
	} else {
		fmt.Println("ES连接成功")
	}

	//设置索引
	IndexName = fmt.Sprintf("gomessage-%s", time.Now().Format("2006.01.02"))
	ctx := context.Background()

	//判断索引是否存在，如果不存在则创建
	exists, err := ElasticsearchClient.IndexExists(IndexName).Do(ctx)
	if err != nil {
		panic(err)
	}

	if !exists {
		//如果索引不存在，则创建一个
		_, err := ElasticsearchClient.CreateIndex(IndexName).BodyString(Mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}
