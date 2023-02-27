package initialize

import (
	"gomessage/assets/docs"
	"gomessage/utils/log/loggers"
)

// InitSwagger 初始化函数
func InitSwagger() {
	docs.SwaggerInfo.Title = "GoMessage"
	docs.SwaggerInfo.Version = "v2.0"
	docs.SwaggerInfo.Description = "承担：消息转发功能；\n\n提供：标准Restful API接口；\n\n支持：同时对多个接收端推送消息；"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = ""

	loggers.DefaultLogger.Info("Swagger模块初始化完成...")
}
