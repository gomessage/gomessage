package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gomessage/authorization"
	"gomessage/pkg/general"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string
		tokenStringList, ok := c.Request.Header["Authorization"]
		if ok {
			tokenString = tokenStringList[0]
			if tokenString == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, general.ResponseFailure("需要Authorization请求头", nil))
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, general.ResponseFailure("需要Authorization请求头", nil))
			return
		}

		//如果token存在于session表中，则代表"登录中"，如果不存在与session表中则代表"没有登陆"
		if len(authorization.QueryToken(tokenString)) >= 1 {
			claims := &authorization.Claims{}

			withClaims, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
				}
				return authorization.JwtKey, nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					c.AbortWithStatusJSON(http.StatusUnauthorized, general.ResponseFailure("未授权", err))
					return
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, general.ResponseFailure("令牌异常", err))
				return
			}

			if !withClaims.Valid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, general.ResponseFailure("令牌无效", nil))
				return
			}

			c.Set("uname", claims.Username)

			c.Next()
		} else {
			//如果token在session表中被标记为"不活跃"，则代表这个token已经"注销"了，不能再使用了
			c.AbortWithStatusJSON(http.StatusUnauthorized, general.ResponseFailure("令牌已过期，请重新登陆", nil))
			return
		}
	}
}
