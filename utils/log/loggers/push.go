package loggers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gomessage/utils/log/hooks/es"
	"io"
	"os"
	"path"
)

var PushLogger = logrus.New()

func InitPushLogger() {
	//日志文件位置
	logFile := viper.GetString("log.pushLogFile")

	//确保存放日志文件的目录始终存在
	logPathDir := path.Dir(logFile)                              //返回路径中除去最后一个元素的剩余部分，也就是路径最后一个元素所在的目录
	if err := os.MkdirAll(logPathDir, os.ModePerm); err != nil { //创建目录类似于（mkdir -p /aaa/bbb的效果）
		fmt.Println("创建push日志目录失败：", err)
		os.Exit(3)
	}

	//确保日志文件始终也存在
	if _, err := os.Create(logFile); err != nil { //创建日志文件
		fmt.Println("创建push日志文件失败：", err)
		os.Exit(3)
	}

	//指定日志输出方式
	//writer1 := os.Stdout
	writer2, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0644) //写入到文件中；此处不使用追加写入，每一次都是覆盖式写入，便于每一次运行时都得到的是干净的日志文件
	if err != nil {
		fmt.Println("创建文件日志文件失败: ", err)
	}

	//设定日志输出位置
	//AccessLogger.SetOutput(io.MultiWriter(writer1, writer2))
	AccessLogger.SetOutput(io.MultiWriter(writer2))

	//设定输出日志中是否要携带上文件名与行号
	AccessLogger.SetReportCaller(false)

	//设定日志等级
	AccessLogger.SetLevel(logrus.InfoLevel)

	//设定日志输出格式
	AccessLogger.SetFormatter(
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05.000 -0700 MST",
		},
	)
	//如果开启es的日志投放功能，则加载对应的钩子
	if viper.GetBool("log.log2es") {
		AccessLogger.AddHook(es.NewPushToEsHook())
	}
}
