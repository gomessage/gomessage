package client

import (
	"github.com/gin-gonic/gin"
	"gomessage/pkg/models"
	"gomessage/pkg/utils"
	"net/http"
)

// ListClient
// @Tags Client
// @Summary 获取所有客户端
// @Router /api/v1/:namespace/client [GET]
func ListClient(g *gin.Context) {
	ns := g.Param("namespace")
	result, err := models.ListClient(ns)
	if err != nil {
		g.JSON(http.StatusInternalServerError, utils.ResponseFailure("查询失败", err))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("获取成功", result))
}
