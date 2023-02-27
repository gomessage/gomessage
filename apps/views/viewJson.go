package views

import (
	"github.com/gin-gonic/gin"
	"gomessage/apps/controllers/hijacking"
	"gomessage/apps/views/httpBase"
	"net/http"
)

func GetNamespaceJson(g *gin.Context) {
	// 从缓存中获取"指定命名空间"中最新一条的数据，然后交给前端进行格式化且展示出来
	ns := g.Param("namespace")
	result, _ := hijacking.GetCacheData(ns)
	g.JSON(http.StatusOK, httpBase.ResponseSuccessful("劫持数据成功", result))
}
