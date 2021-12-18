package controllers

import beego "github.com/beego/beego/v2/server/web"

//首页
type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "beego@gmail.com"
	c.TplName = "index.html"
}
