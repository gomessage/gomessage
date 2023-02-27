package httpHealth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health
// @Tags Health
// @Summary 全局健康检测
// @Router /health [GET]
func Health(g *gin.Context) {
	g.JSON(http.StatusOK, "Health monitoring is normal")
}

// Hello
// @Tags Health
// @Summary 命名空间健康检测
// @Description 访问该接口将返回：Hello World
// @Accept json
// @Produce json
// @Success 200 {string} {string}
// @Router /api/v1/:namespace/health [GET]
func Hello(g *gin.Context) {
	g.JSON(http.StatusOK, "Health monitoring is normal")
}
