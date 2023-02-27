package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomessage/routers"
	"gomessage/utils/initialize"
	"gomessage/utils/log/loggers"
)

func init() {
	/*
	 * **********************************************
	 *
	 * 特别注意：初始化其他模块时，需要严格注意启动的先后顺序
	 *
	 * initConfig() 应该被第一个执行
	 *
	 * initEnv() 应该被第二个执行（这里会接收一些命令行启动参数，会覆盖掉InitConfig中的一部分参数，这样开发调试时就不用频繁修改配置文件了）
	 *
	 * initLog() 应该被第三个执行（剩下的其它组件就没啥先后顺序要求了）
	 *
	 * **********************************************
	 */
	initialize.InitConfig() //应该被第一个执行
	initialize.InitEnv()    //初始化环境变量，应该紧跟在InitConfig后面，被第二个执行
	loggers.InitLog()       //初始化日志模块，应该紧跟在InitEnv后面，被第三个执行（剩下的其它模块初始化，启动顺序就没有什么要求了）
	initialize.InitSwagger()
	initialize.InitGinMode()
	initialize.InitDB("sqlite3", initialize.GlobalVars.Migrate)
}

func main() {
	//创建gin实例
	r := gin.Default()

	//加载路由表
	routers.Path(r)

	//捕获全局panic的错误
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	//启动web服务，默认(BaseURL := "127.0.0.1:7077")
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	loggers.DefaultLogger.Info("服务监听：http://127.0.0.1:7077")
	if err := r.Run(":7077"); err != nil {
		return
	}
}
