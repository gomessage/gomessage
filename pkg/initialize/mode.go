package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gomessage/pkg/log/loggers"
)

// InitGinMode 初始化函数
func InitGinMode() {
	//设定gin的运行模式，ginMode的值，后面可以考虑从viper中获取。
	ginMode := viper.GetString("global.ginMode")
	switch ginMode {
	case "debug":
		gin.SetMode(gin.DebugMode)

	case "test":
		gin.SetMode(gin.TestMode)

	case "release":
		gin.SetMode(gin.ReleaseMode)

	default:
		//默认为"release"
		gin.SetMode(gin.ReleaseMode)
	}
	loggers.DefaultLogger.Info("RunMode初始化完成...")
}
