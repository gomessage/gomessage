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
	//首页
	beego.Router("/", &controllers.IndexController{})

	//多管道~消息入口（测试版）
	beego.Router("/gomessage", &web2.PipelineControllers{})
	beego.Router("/gomessage/:namespace:string", &web2.PipelineControllers{})

	//接口：v1版本的api
	ns := beego.NewNamespace("/v1",
		//一个命名空间(Namespace)下的多个模块拼装在一起，才能形成一个"管道（或消息通道）"的概念~
		//前端对用户传达"管道"的概念或UI名称，是为了便于用户从"使用层面"快速的理解GoMessage"，但是后端开发过程中要意识到"命名空间"存在的重要意义~
		//未来的某段时间内，如果您也参与进来为GoMessage贡献代码，请尝试着把Namespace视为一个重要的"逻辑平面"或某种"主键"；
		//从数据库表设计、到抽象方法的封装、到业务逻辑的编写、到接口的设计，每一层都把Namespace视作一个整体，从而完整的构建出了GoMessage的"多通道"处理模型~

		//命名空间管理路由
		beego.NSNamespace("/namespace",
			beego.NSInclude(&web2.NamespaceControllers{}),
		),
		//用户变量映射路由
		beego.NSNamespace("/map",
			beego.NSInclude(&web2.MapControllers{}),
		),
		//模板编辑路由
		beego.NSNamespace("/template",
			beego.NSInclude(&web2.TemplateControllers{}),
		),
		//客户端操作路由
		beego.NSNamespace("/client",
			beego.NSInclude(&web2.Clients{}),
			beego.NSInclude(&web2.Client{}),
			beego.NSInclude(&web2.ClientActive{}),
		),
		//数据劫持路由
		beego.NSNamespace("/metadata",
			beego.NSInclude(&web2.JsonControllers{}),
		),
		//测试接口
		//beego.NSNamespace("/test",
		//	beego.NSInclude(&test.TestController{}),
		//),
	)

	//注册命名空间
	beego.AddNamespace(ns)
	//beego.AddNamespace(ns2)
}
