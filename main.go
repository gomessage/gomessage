package main

import (
	_ "GoMessage/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/api"] = "swagger"
	}
	beego.Run()
}
