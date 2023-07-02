package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	api2 "gomessage/src/api"
	client2 "gomessage/src/api/client"
	"gomessage/src/authorization"
	middleware2 "gomessage/src/middleware"
	"gomessage/src/pkg/log/loggers"
	"net/http"
)

func AddStatic(g *gin.Engine) {
	g.StaticFile("/favicon.ico", "./assets/favicon.ico")
	g.Static("/static", "assets/vue/static")
	g.LoadHTMLGlob("assets/vue/*.html")
}

func Path(g *gin.Engine) {

	//加载静态文件
	AddStatic(g)

	//中间件
	g.Use(middleware2.Cors())
	g.Use(middleware2.AccessLog())

	//健康检测
	g.GET("/ok", api2.Hello)
	g.GET("/health", api2.Health)

	//首页
	g.GET("/", api2.Index)

	//swagger文档
	g.GET("/docs", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "/swagger/index.html") })
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //Swagger页面

	//数据入口Get方法
	g.GET("/go/:namespace", middleware2.CheckNamespace(), api2.GoMessageByGet)                           //给单个路由追加中间件middleware.CheckNamespace()
	g.GET("/go", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) })        //把"/go"重定向到"/go/message"的路由上
	g.GET("/gomessage", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) }) //把"/gomessage"重定向到"/go/message"的路由上
	//数据入口Post方法
	g.POST("/go/:namespace", middleware2.CheckNamespace(), api2.GoMessageByTransport)                     //给单个路由追加中间件middleware.CheckNamespace()
	g.POST("/go", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) })        //把"/go"重定向到"/go/message"的路由上
	g.POST("/gomessage", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) }) //把"/gomessage"重定向到"/go/message"的路由上

	//登录注册相关
	g.POST("/auth/login", authorization.Login)
	g.POST("/auth/logout", authorization.Logout)
	g.POST("/auth/register", authorization.Register)

	//=======================
	// 用户操作：v1版本
	//=======================
	v1View := g.Group("/api/v1")
	v1View.Use(middleware2.CheckNamespace())
	v1View.Use(middleware2.AuthMiddleware())
	{
		//命名空间健康检测
		v1View.GET("/:namespace/health", api2.Health)

		//数据劫持
		v1View.GET("/:namespace/json", api2.GetNamespaceJson)

		//用户变量
		v1View.GET("/:namespace/vars", api2.ListVariables)  //获取变量映射
		v1View.POST("/:namespace/vars", api2.PostVariables) //添加变量映射

		//TODO: 这个功能虽然后端完成了，但是前端ui层面没有启动这个接口对应的功能，有空时再过来修改一下
		v1View.GET("/:namespace/flattening", api2.GetNamespaceFlatteningJson)

		//消息模板
		v1View.GET("/:namespace/template", api2.ListTemplate)
		v1View.POST("/:namespace/template", api2.PostTemplate)

		//客户端
		v1View.GET("/:namespace/client", client2.ListClient)             //获取所有客户端
		v1View.POST("/:namespace/client", client2.PostClient)            //新增一个客户端
		v1View.GET("/:namespace/client/:id", client2.GetClient)          //获取客户端详情
		v1View.PUT("/:namespace/client/:id", client2.PutClientActive)    //更新一个客户端
		v1View.PUT("/:namespace/client-info/:id", client2.PutClientInfo) //更新一个客户端
		v1View.DELETE("/:namespace/client/:id", client2.DeleteClient)    //删除一个客户端

	}

	//=======================
	// 命名空间操作：v1版本
	//=======================
	v1Namespace := g.Group("/api/v1")
	v1Namespace.Use(middleware2.AuthMiddleware())
	{
		//命名空间
		v1Namespace.GET("/namespace", api2.ListNamespace)
		v1Namespace.POST("/namespace", api2.PostNamespace)
		v1Namespace.GET("/namespace/:id", api2.GetNamespace)
		v1Namespace.PUT("/namespace/:id", api2.PutNamespace)
		v1Namespace.DELETE("/namespace/:id", api2.DeleteNamespace) //删除命名空间的时候，需要把当前命名空间下的（变量映射、模板、客户端）全都删除掉
	}

	loggers.DefaultLogger.Info("路由表加载完成...")

}
