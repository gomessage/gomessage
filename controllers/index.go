package controllers

import beego "github.com/beego/beego/v2/server/web"

//首页
type IndexController struct {
	beego.Controller
}

//GoMessage首页，用来显示vue.js的静态页面
func (c *IndexController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "beego@gmail.com"
	c.TplName = "index.html"
}
