package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// InitConfig 初始化函数（在main函数里面被引用，全局第一个被加载的函数）
func InitConfig() {

	//默认配置文件路径（以后可以被其它形式替换掉，这样就可以支持多配置文件了）
	defaultConfigPath := "config/default.yaml"

	//加载的配置文件（写法二）
	viper.SetConfigFile(defaultConfigPath)

	//读取配置文件（读取上面指定的配置文件）
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("启动日志：基础"+defaultConfigPath+"配置文件读取失败：", err)
		os.Exit(3)
	} else {
		fmt.Println("启动日志：基础" + defaultConfigPath + "配置文件读取完成...")
	}

	//初始化Config：这里是给一些运行参数赋予默认值，如果配置文件中没有填写，且启动参数也没有指定，那就加载这里的默认值
	setDefaultBaseConfig()

	/*
	 * **********************************************
	 *
	 * 特别注意：初始化其他模块时，需要严格注意启动的先后顺序
	 *
	 * initConfig 应该被第一个执行（也就是此函数本身）
	 *
	 * initEnv 应该被第二个执行（这里的启动参数，会覆盖掉InitConfig中的一部分参数，这样开发调试时就不用频繁修改配置文件了）
	 *
	 * initLog 应该被第三个执行（剩下的其它组件就没啥先后顺序要求了）
	 *
	 * **********************************************
	 */
	//InitConfig()	   						     //这就是当前函数自身，注释在这里只是为了提醒顺序
	//InitEnv()                                  //初始化环境变量，应该紧跟在InitConfig后面，被第二个执行
	//runLog.InitLog()                           //初始化日志模块，应该紧跟在InitEnv后面，被第三个执行（剩下的其它模块初始化，启动顺序就没有什么要求了）
	//InitSwagger()                              //初始化Swagger
	//InitGinMode()                              //初始化GinMode
	//InitRunParams()                            //初始化运行时参数
	//InitDB(GlobalVars.Env, GlobalVars.Migrate) //初始化数据库

}

// 设置默认配置项
func setDefaultBaseConfig() {
	/*
	   如果 config/default.yaml 中没有设置下面这些内容，
	   则下文中的默认值将会生效；
	   配置文件中的优先级高于此处设定的优先级
	*/

	//全局相关
	viper.SetDefault("global.env", "fat")         //当前运行环境
	viper.SetDefault("global.ginMode", "release") //当前gin的运行模式；（选项：debug、release）
	viper.SetDefault("databaseType", "sqlite3")   //当前连接数据库类型

	//运行相关
	viper.SetDefault("server.host", "127.0.0.1") //启动时监听的地址
	viper.SetDefault("server.port", "8080")      //启动时监听的端口

	//日志相关
	viper.SetDefault("log.level", "info")              //日志记录级别；（选项：debug、info）
	viper.SetDefault("log.format", "json")             //日志记录格式；（选项：json、text）
	viper.SetDefault("log.path", "./logs/service.log") //日志输出位置
}
