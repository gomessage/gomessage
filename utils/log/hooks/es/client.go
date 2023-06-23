package es

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
	"time"
)

/*
 *
 * 如果需要一个开关，来控制日志写入es的功能是否启用，那么这个开关应该控制两个地方：
 * - es初始化的位置
 * - es钩子是否被加载的位置
 *
 */

var elasticsearchClient *elastic.Client
var once sync.Once

func GetEsClient() *elastic.Client {
	once.Do(func() {
		if viper.GetBool("log.log2es") {
			//如果log.log2es==true，则开启向es的投放日志，初始化es客户端
			initClient()
		}
	})
	return elasticsearchClient
}

func initClient() {
	//var esUrl = []string{"http://xxx.xxx.xxx:9200", "http://xxx.xxx.xxx:9200", "http://xxx.xxx.xxx:9200"}
	var esUrl = viper.GetStringSlice("log.es.urls")
	var esUserName = viper.GetString("log.es.username")
	var esPassword = viper.GetString("log.es.password")
	var err error

	/*
	 * TODO 如果本地环境变量存在以下内容，使用以下地址覆盖之前的地址，获得最高优先级，便于本地开发调试.
	 * TODO 此处的EsUrl在全局环境变量声明时，只是一个普通的字符串类型，如果需要兼容多个地址，可以自行修改此处代码
	 */
	envEsUrl := []string{os.Getenv("EsUrl")}
	envEsUsername := os.Getenv("EsUsername")
	envEsPassword := os.Getenv("EsPassword")
	if len(envEsUrl) > 0 && envEsUsername != "" && envEsPassword != "" {
		esUrl = envEsUrl
		esUserName = envEsUsername
		esPassword = envEsPassword
	}

	elasticsearchClient, err = elastic.NewClient(
		elastic.SetURL(esUrl...),
		elastic.SetBasicAuth(esUserName, esPassword),
		elastic.SetGzip(true),
		elastic.SetHealthcheckInterval(2*time.Second),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC：", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)), //平时调试的时候开启，发行版中不用开启
		elastic.SetSniff(false),
	)

	if err != nil {
		fmt.Println("ES连接失败：", err)
	} else {
		fmt.Println("ES连接成功")
	}

	//创建多个索引
	for _, indexName := range []string{GetIndexName("gomessage-access"), GetIndexName("gomessage-runtime"), GetIndexName("gomessage-push")} {
		SyncIndex(indexName)
	}
}
