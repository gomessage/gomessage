package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gomessage/api"
	"gomessage/api/client"
	"gomessage/middleware"
	"gomessage/pkg/log/loggers"
	"net/http"
)

func initStatic(g *gin.Engine) {
	g.StaticFile("/favicon.ico", "./assets/favicon.ico")
	g.Static("/static", "assets/vue/static")
	g.LoadHTMLGlob("assets/vue/*.html")
}

func Path(g *gin.Engine) {

	//=======================
	// 全局路由：基础部分
	//=======================
	//中间件
	g.Use(middleware.Cors())
	g.Use(middleware.AccessLog())

	// Once it's done, you can attach the handler as one of your middleware
	//g.Use(sentrygin.New(sentrygin.Options{}))
	//g.Use(func(ctx *gin.Context) {
	//	if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
	//		hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
	//	}
	//	ctx.Next()
	//})

	//加载静态文件
	initStatic(g)
	//路由重定向
	g.GET("/", api.Index)
	g.GET("/docs", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "/swagger/index.html") })
	//基础URI
	g.GET("/ok", api.Health)                                             //健康监测
	g.GET("/health", api.Health)                                         //健康监测
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //Swagger页面

	//=======================
	// gomessage数据入口
	//=======================
	g.GET("/go/:namespace", middleware.IsNamespace(), api.GoMessageByGet)                                //给单个路由追加中间件middleware.IsNamespace()
	g.GET("/go", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) })        //把"/go"重定向到"/go/message"的路由上
	g.GET("/gomessage", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) }) //把"/gomessage"重定向到"/go/message"的路由上
	//接收数据推送
	g.POST("/go/:namespace", middleware.IsNamespace(), api.GoMessageByTransport)                          //给单个路由追加中间件middleware.IsNamespace()
	g.POST("/go", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) })        //把"/go"重定向到"/go/message"的路由上
	g.POST("/gomessage", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) }) //把"/gomessage"重定向到"/go/message"的路由上

	//=======================
	// 用户操作接口：v1版本
	//=======================
	v1View := g.Group("/api/v1")
	v1View.Use(middleware.IsNamespace())
	{
		//命名空间健康检测
		v1View.GET("/:namespace/health", api.Health)

		//数据劫持
		v1View.GET("/:namespace/json", api.GetNamespaceJson)

		//用户变量
		v1View.GET("/:namespace/vars", api.ListVariables)  //获取变量映射
		v1View.POST("/:namespace/vars", api.PostVariables) //添加变量映射
		//TODO: 这个功能虽然后端完成了，但是前端ui层面没有启动这个接口对应的功能，有空时再过来修改一下
		v1View.GET("/:namespace/flattening", api.GetNamespaceFlatteningJson)

		//消息模板
		v1View.GET("/:namespace/template", api.ListTemplate)
		v1View.POST("/:namespace/template", api.PostTemplate)

		//客户端
		v1View.GET("/:namespace/client", client.ListClient)          //获取所有客户端
		v1View.POST("/:namespace/client", client.PostClient)         //新增一个客户端
		v1View.GET("/:namespace/client/:id", client.GetClient)       //获取客户端详情
		v1View.PUT("/:namespace/client/:id", client.PutClient)       //更新一个客户端
		v1View.DELETE("/:namespace/client/:id", client.DeleteClient) //删除一个客户端

	}

	//=======================
	// 命名空间操作接口：v1版本
	//=======================
	v1Namespace := g.Group("/api/v1")
	{
		//命名空间
		v1Namespace.GET("/namespace", api.ListNamespace)
		v1Namespace.POST("/namespace", api.PostNamespace)
		v1Namespace.GET("/namespace/:id", api.GetNamespace)
		v1Namespace.PUT("/namespace/:id", api.PutNamespace)
		v1Namespace.DELETE("/namespace/:id", api.DeleteNamespace) //删除命名空间的时候，需要把当前命名空间下的（变量映射、模板、客户端）全都删除掉
	}

	loggers.DefaultLogger.Info("路由表加载完成...")

}
