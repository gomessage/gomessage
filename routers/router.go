// @APIVersion v1
// @Title GoMessage
// @Description 转发Json格式的信息，到钉钉、微信、或其它客户端
// @Contact tay3223@qq.com
// //@TermsOfServiceUrl https://blog.taycc.com/pages/prometheus/target_icmp.html
// //@License Apache 2.0
// //@LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"GoMessage/controllers"
	"GoMessage/controllers/web2"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//ns2 := beego.NewNamespace("/test",
	//	//与web界面的api没有任何关系，唯一有关系的可能是web2里有些函数引用了老的结构体
	//	beego.NSNamespace("/alertmanager",
	//		beego.NSInclude(&alertmanager.K8sControllers{}),
	//		beego.NSInclude(&alertmanager.LinuxControllers{}),
	//	),
	//
	//	//测试接口
	//	beego.NSNamespace("/test",
	//		beego.NSInclude(&test.TestController{}),
	//	),
	//)

	//首页~前端静态页面
	beego.Router("/", &controllers.IndexController{})

	//单管道~消息入口
	beego.Router("/go/message", &web2.ApiControllers{})

	//多管道~消息入口（测试版）
	beego.Router("/gomessage/:label:string", &web2.PipelineControllers{})

	//命名空间（v1版本的api）
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/namespace",
			beego.NSInclude(&web2.NamespaceControllers{}),
		),
		beego.NSNamespace("/map",
			beego.NSInclude(&web2.MapControllers{}),
		),
		beego.NSNamespace("/template",
			beego.NSInclude(&web2.TemplateControllers{}),
		),
		beego.NSNamespace("/client",
			beego.NSInclude(&web2.Clients{}),
			beego.NSInclude(&web2.Client{}),
			beego.NSInclude(&web2.ClientActive{}),
		),
		//GoMessage业务的api
		beego.NSNamespace("/metadata",
			beego.NSInclude(&web2.JsonControllers{}),

			//beego.NSRouter("/client", &web2.Clients{}),
			//beego.NSRouter("/client/:id:int", &web2.Client{}),
			//beego.NSRouter("/client/active", &web2.ClientActive{}),
		),
	)

	//注册命名空间
	beego.AddNamespace(ns)
	//beego.AddNamespace(ns2)
}
