package client

import (
	"github.com/gin-gonic/gin"
	"gomessage/api"
	"gomessage/models"
	"net/http"
)

// ListClient
// @Tags Client
// @Summary 获取所有客户端
// @Router /api/v1/:namespace/client [GET]
func ListClient(g *gin.Context) {
	ns := g.Param("namespace")
	result, _ := models.ListClient(ns)
	g.JSON(http.StatusOK, api.ResponseSuccessful("获取成功", result))
}
