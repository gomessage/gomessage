package runLog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
	"path"
)

// Log 此处实例化一个日志对象，是用来记录服务自身'运行日志'的，request请求过程中的access日志由另外一个封装好的'日志中间件'来单独记录
var Log = logrus.New()

// InitLog 初始化函数
func InitLog() {

	//日志文件位置
	logFile := viper.GetString("log.logFile") //获取默认配置文件中指定的值
	//logFile := "./logs/service.log"              //获取默认配置文件中指定的值

	//确保存放日志文件的目录始终存在
	logPathDir := path.Dir(logFile)                              //返回路径中除去最后一个元素的剩余部分，也就是路径最后一个元素所在的目录
	if err := os.MkdirAll(logPathDir, os.ModePerm); err != nil { //创建目录类似于（mkdir -p /aaa/bbb的效果）
		fmt.Println("创建日志目录失败：", err)
		os.Exit(3)
	}

	//确保日志文件始终也存在
	if _, err := os.Create(logFile); err != nil { //创建日志文件
		fmt.Println("创建日志文件失败：", err)
		os.Exit(3)
	}

	//指定日志输出方式
	writer1 := os.Stdout                                                //标准输出
	writer2, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0644) //写入到文件中；此处不使用追加写入，每一次都是覆盖式写入，便于每一次运行时都得到的是干净的日志文件
	if err != nil {
		fmt.Println("创建文件日志文件失败: ", err)
	}

	//设定日志输出位置
	Log.SetOutput(io.MultiWriter(writer1, writer2))

	//设定输出日志中是否要携带上文件名与行号
	Log.SetReportCaller(false)

	//设定日志等级
	setLogLevel(Log, viper.GetString("log.level"))

	//设定日志输出格式
	setLogFormat(Log, viper.GetString("log.format"))

	Log.Info("log模块初始化完成...")
}

// 设定日志等级
func setLogLevel(log *logrus.Logger, level string) {
	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)

	case "info":
		log.SetLevel(logrus.InfoLevel)

	default:
		log.SetLevel(logrus.DebugLevel)
	}
	/*
	   logrus有一个日志级别，默认的级别为InfoLevel。如果为了能看到Trace和Debug日志，需要在设置日志级别为TraceLevel。
	   Log.Trace("trace msg") //很细粒度的信息，一般用不到
	   Log.Debug("debug msg") //一般程序中输出的调试信息
	   Log.Info("info msg")   //关进操作，核心流程的日志
	   Log.Warn("warn msg")   //警告信息，提醒程序员注意
	   Log.Error("error msg") //错误日志，需要查看原因
	   Log.Panic("panic msg") //记录日志，然后panic
	   Log.Fatal("fatal msg") //致命错误，出现错误时，程序无法正常运转，输出日志后，程序退出
	*/
}

// 设定日志格式
func setLogFormat(log *logrus.Logger, format string) {
	//指定日志输出格式
	switch format {
	case "json":
		//设置日志的输出格式（json格式）
		log.SetFormatter(
			&logrus.JSONFormatter{
				TimestampFormat: "2006-01-02 15:04:05.000 -0700 MST",
			},
		)
	case "text":
		//设置日志的输出格式（Text格式）
		log.SetFormatter(
			&logrus.TextFormatter{
				TimestampFormat: "2006-01-02 15:04:05.000 -0700 MST",
			},
		)
	default:
		//设置日志的输出格式（Text格式）
		log.SetFormatter(
			&logrus.TextFormatter{
				TimestampFormat: "2006-01-02 15:04:05.000 -0700 MST",
			},
		)
	}
}
