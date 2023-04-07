package httpIndex

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index
// Vue静态资源的首页，需要先在路由表中加载如下两个内容：
// g.Static("/static", "assets/vue/static")
// g.LoadHTMLGlob("assets/vue/*.html")
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "GoMessage"})
}
