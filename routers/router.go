// @APIVersion 1.0.0
// @Title 信息转发器
// @Description 转发Prometheus生态圈中AlertManager推送的webhook信息到钉钉或微信等
// @Contact tay3223@11.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"GoMessage/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/monitoring",
		beego.NSNamespace("/object",
			beego.NSInclude(&controllers.ObjectController{}),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.UserController{}),
		),
		beego.NSNamespace("/prometheus",
			beego.NSInclude(&controllers.DingtalkControllers{}),
		),
	)
	beego.AddNamespace(ns)
}
