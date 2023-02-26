package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
        c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
        c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
        c.Header("Access-Control-Allow-Credentials", "true")

        // 放行所有OPTIONS方法，因为有的模板是要请求两次的
        if method == "OPTIONS" {
            c.AbortWithStatusJSON(http.StatusOK, "ok")
        }

        // 处理请求
        c.Next()
    }
}
