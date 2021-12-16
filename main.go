package main

import (
	_ "GoMessage/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	//日志级别
	//logs.SetLevel(logs.LevelDebug)
	//logs.SetLogger("console")
	//logs.SetLogger(logs.AdapterFile, `{"filename":"app.log", "level":9}`)

	//连接数据库
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "conf/db.sqlite3")
	orm.RunSyncdb("default", false, true)
	//orm.Debug = true //开启调试模式
}

func main() {
	//解决跨域问题
	//InsertFilter是提供一个过滤函数
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))

	if beego.BConfig.RunMode == "dev" || beego.BConfig.RunMode == "test" || beego.BConfig.RunMode == "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/api"] = "swagger"
	}
	beego.Run()
}
