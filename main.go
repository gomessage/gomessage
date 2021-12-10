package main

import (
	_ "GoMessage/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//日志级别
	//logs.SetLevel(logs.LevelDebug)
	//logs.SetLogger("console")
	//logs.SetLogger(logs.AdapterFile, `{"filename":"app.log", "level":9}`)
}

func main() {
	if beego.BConfig.RunMode == "dev" || beego.BConfig.RunMode == "test" || beego.BConfig.RunMode == "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/"] = "swagger"
	}
	beego.Run()
}
