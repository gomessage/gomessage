// @APIVersion 1.0.0
// @Title 信息转发器
// @Description 转发Alertmanager的webhook信息，到钉钉、微信、或其它客户端
// @Contact tay3223@qq.com
// @TermsOfServiceUrl https://blog.taycc.com/pages/prometheus/target_icmp.html
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"GoMessage/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/client",
			beego.NSInclude(&controllers.DingtalkControllers{}),
		),
	)
	beego.AddNamespace(ns)
}
