package routers

import (
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "gomessage/apps/views"
    "gomessage/apps/views/httpClient"
    "gomessage/apps/views/httpHealth"
    "gomessage/apps/views/httpIndex"
    "gomessage/routers/middleware"
    "gomessage/utils/log/loggers"
    "net/http"
)

func initStatic(g *gin.Engine) {
    g.StaticFile("/favicon.ico", "./assets/favicon.ico")
    g.Static("/static", "assets/vue2/static")
    g.LoadHTMLGlob("assets/vue2/*.html")
}

func Path(g *gin.Engine) {

    //=======================
    // 全局路由：基础部分
    //=======================
    //中间件
    g.Use(middleware.Cors())
    g.Use(middleware.AccessLog())
    //加载静态文件
    initStatic(g)
    //路由重定向
    g.GET("/", httpIndex.Index)
    g.GET("/docs", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "/swagger/index.html") })
    //基础URI
    g.GET("/ok", httpHealth.Health)                                      //健康监测
    g.GET("/health", httpHealth.Health)                                  //健康监测
    g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //Swagger页面

    //=======================
    // gomessage数据入口
    //=======================
    g.GET("/go/:namespace", middleware.IsNamespace(), views.GoMessageByGet)                              //给单个路由追加中间件middleware.IsNamespace()
    g.GET("/go", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) })        //把"/go"重定向到"/go/message"的路由上
    g.GET("/gomessage", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) }) //把"/gomessage"重定向到"/go/message"的路由上
    //接收数据推送
    g.POST("/go/:namespace", middleware.IsNamespace(), views.GoMessageByPost)                             //给单个路由追加中间件middleware.IsNamespace()
    g.POST("/go", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) })        //把"/go"重定向到"/go/message"的路由上
    g.POST("/gomessage", func(c *gin.Context) { c.Request.URL.Path = "/go/message"; g.HandleContext(c) }) //把"/gomessage"重定向到"/go/message"的路由上

    //=======================
    // 用户操作接口：v1版本
    //=======================
    v1View := g.Group("/api/v1")
    v1View.Use(middleware.IsNamespace())
    {
        //命名空间健康检测
        v1View.GET("/:namespace/health", httpHealth.Health)

        //数据劫持
        v1View.GET("/:namespace/json", views.GetNamespaceJson)

        //客户端
        v1View.GET("/:namespace/client", httpClient.ListClient)          //获取所有客户端
        v1View.POST("/:namespace/client", httpClient.PostClient)         //新增一个客户端
        v1View.GET("/:namespace/client/:id", httpClient.GetClient)       //获取一个客户端的详情
        v1View.PUT("/:namespace/client/:id", httpClient.PutClient)       //更新一个客户端
        v1View.DELETE("/:namespace/client/:id", httpClient.DeleteClient) //删除一个客户端

        //用户变量
        v1View.GET("/:namespace/vars", views.ListVariables)
        v1View.POST("/:namespace/vars", views.PostVariables)

        //消息模板
        v1View.GET("/:namespace/template", views.ListTemplate)
        v1View.POST("/:namespace/template", views.PostTemplate)

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

    loggers.DefaultLogger.Info("路由表加载完成...")

}
