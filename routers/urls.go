package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gomessage/apps/views"
	"gomessage/apps/views/health"
	"gomessage/apps/views/index"
	"gomessage/apps/views/viewClient"
	"gomessage/routers/middleware"
	"gomessage/utils/runLog"
	"net/http"
)

func initStatic(g *gin.Engine) {
	g.StaticFile("/favicon.ico", "./assets/favicon.ico")
	g.Static("/static", "assets/vue2/static")
	g.LoadHTMLGlob("assets/vue2/*.html")
}

func Path(g *gin.Engine) {

	//=======================
	// 全局基础路由
	//=======================
	//中间件
	g.Use(middleware.Cors())
	//加载静态文件
	initStatic(g)

	//路由重定向
	g.GET("/", index.Index)
	g.GET("/docs", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "/swagger/index.html") })

	//基础URI
	g.GET("/ok", health.Health)                                          //健康监测
	g.GET("/health", health.Health)                                      //健康监测
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //Swagger页面

	//=======================
	// /go/message数据入口
	//=======================
	g.GET("/go/:namespace", views.GoMessageByGet, middleware.IsNamespace())
	g.GET("/go", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) })        //把几个路由关键字，统一重定向到统一的处理逻辑上
	g.GET("/gomessage", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) }) //把几个路由关键字，统一重定向到统一的处理逻辑上

	g.POST("/go/:namespace", views.GoMessageByPost, middleware.IsNamespace())
	g.POST("/go", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) })        //把几个路由关键字，统一重定向到统一的处理逻辑上
	g.POST("/gomessage", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) }) //把几个路由关键字，统一重定向到统一的处理逻辑上

	//=======================
	// 用户操作接口：v1版本
	//=======================
	v1View := g.Group("/api/v1")
	v1View.Use(middleware.IsNamespace())
	{
		//命名空间健康检测
		v1View.GET("/:namespace/health", health.Health)

		//数据劫持
		v1View.GET("/:namespace/json", views.GetNamespaceJson)

		//客户端
		v1View.GET("/:namespace/client", viewClient.ListClient)          //获取所有客户端
		v1View.POST("/:namespace/client", viewClient.PostClient)         //新增一个客户端
		v1View.GET("/:namespace/client/:id", viewClient.GetClient)       //获取一个客户端的详情
		v1View.PUT("/:namespace/client/:id", viewClient.PutClient)       //更新一个客户端
		v1View.DELETE("/:namespace/client/:id", viewClient.DeleteClient) //删除一个客户端

		//用户变量
		v1View.GET("/:namespace/vars", views.ListVariables)
		v1View.POST("/:namespace/vars", views.PostVariables)
		//v1View.GET("/:namespace/vars/:id", views.GetVariables)
		//v1View.PUT("/:namespace/vars/:id", views.PutVariables)
		//v1View.DELETE("/:namespace/vars/:id", views.DeleteVariables)

		//消息模板
		v1View.GET("/:namespace/template", views.ListTemplate)
		v1View.POST("/:namespace/template", views.PostTemplate)
		//v1View.GET("/:namespace/template/:id", views.GetTemplate)
		//v1View.PUT("/:namespace/template/:id", views.PutTemplate)
		//v1View.DELETE("/:namespace/template/:id", views.DeleteTemplate)
	}

	//=======================
	// 命名空间操作接口：v1版本
	//=======================
	v1Namespace := g.Group("/api/v1")
	{
		//命名空间
		v1Namespace.GET("/namespace", views.ListNamespace)
		v1Namespace.POST("/namespace", views.PostNamespace)
		v1Namespace.GET("/namespace/:id", views.GetNamespace)
		v1Namespace.PUT("/namespace/:id", views.PutNamespace)
		v1Namespace.DELETE("/namespace/:id", views.DeleteNamespace)
	}

	runLog.Log.Info("路由表加载完成...")

}
