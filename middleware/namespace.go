package middleware

import (
	"github.com/gin-gonic/gin"
	"gomessage/models"
	"gomessage/pkg/utils"
	"net/http"
	"time"
)

type BaseResponse struct {
	Msg       string    `json:"msg"`
	Route     string    `json:"route"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
	ClientIp  string    `json:"client_ip"`
	ErrorTime time.Time `json:"error_time"`
}

// CheckNamespace 中间件：判断Namespace是否存在
func CheckNamespace() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ns string
		if c.Request.URL.Path == "/go" || c.Request.URL.Path == "/gomessage" || c.Request.URL.Path == "/go/message" {
			ns = "default"
		} else {
			ns = c.Param("namespace")
		}

		if !models.IsNamespaceExist(ns) {
			rsp := BaseResponse{
				Msg:       "Path参数 {:namespace} 错误，不存在名为 {" + ns + "} 的命名空间",
				Route:     c.FullPath(),
				Path:      c.Request.URL.Path,
				Method:    c.Request.Method,
				ClientIp:  c.ClientIP(),
				ErrorTime: time.Now(),
			}
			//截断Request请求，直接返回指定状态码和内容
			c.AbortWithStatusJSON(http.StatusNotFound, utils.ResponseFailure("中间件拦截", &rsp))
			return
		} else {
			c.Next()
		}
	}
}
