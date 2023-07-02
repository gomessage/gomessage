package loggers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gomessage/src/pkg/log/hooks/es"
	"io"
	"os"
	"path"
)

var AccessLogger = logrus.New()

func InitAccessLog() {
	//日志文件位置
	logFile := viper.GetString("log.accessLogFile")

	//确保存放日志文件的目录始终存在
	logPathDir := path.Dir(logFile)                              //返回路径中除去最后一个元素的剩余部分，也就是路径最后一个元素所在的目录
	if err := os.MkdirAll(logPathDir, os.ModePerm); err != nil { //创建目录类似于（mkdir -p /aaa/bbb的效果）
		fmt.Println("创建access日志目录失败：", err)
		os.Exit(3)
	}

	//确保日志文件始终也存在
	if _, err := os.Create(logFile); err != nil { //创建日志文件
		fmt.Println("创建access日志文件失败：", err)
		os.Exit(3)
	}

	//指定日志输出方式
	//writer1 := os.Stdout
	writer2, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("创建文件日志文件失败: ", err)
	}

	//设定输出日志中是否要携带上文件名与行号
	AccessLogger.SetReportCaller(false)

	//设定日志等级
	AccessLogger.SetLevel(logrus.InfoLevel)

	//设定日志输出格式
	AccessLogger.SetFormatter(
		&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000 -0700 MST",
		},
	)
	//如果开启es的日志投放功能，则加载对应的钩子
	if viper.GetBool("log.log2es") {
		AccessLogger.AddHook(es.NewAccessToEsHook())
	}

	//设定日志输出位置
	AccessLogger.SetOutput(io.MultiWriter(writer2))
}
