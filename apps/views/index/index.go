package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index
// Vue静态资源的首页，需要先在路由表中加载如下两个内容：
// g.Static("/", "assets/static")
// g.LoadHTMLGlob("assets/static/*.html")
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "GoMessage"})
}
