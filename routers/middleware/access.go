package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gomessage/utils/log/loggers"
	"io"
	"time"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		//开始时间
		startTime := time.Now()

		//uri路径
		uriPath := c.Request.URL.Path

		//请求参数
		raw := c.Request.URL.RawQuery

		if raw != "" {
			uriPath = uriPath + "?" + raw
		}

		// 请求体是一个readClose，它只能被读取一次。
		var body []byte
		if c.Request != nil && c.Request.Body != nil {
			var buf bytes.Buffer
			tee := io.TeeReader(c.Request.Body, &buf)
			body, _ = io.ReadAll(tee)
			c.Request.Body = io.NopCloser(&buf)
		}

		//处理请求
		c.Next()

		//结束时间
		endTime := time.Now()

		//时间偏移量
		latency := endTime.Sub(startTime)

		//记录日志
		loggers.AccessLogger.WithFields(logrus.Fields{
			"start_time":   startTime.Format("2006-01-02_15:04:05.000000"), //开始时间
			"end_time":     endTime.Format("2006-01-02_15:04:05.000000"),   //结束时间
			"latency":      latency.String(),                               //延迟
			"status_code":  c.Writer.Status(),                              //状态码
			"client_ip":    c.ClientIP(),                                   //客户端ip
			"method":       c.Request.Method,                               //请求方法
			"path":         uriPath,                                        //uri路径
			"router":       c.FullPath(),                                   //路由
			"request_body": string(body),                                   //请求数据
		}).Info("access日志记录")

	}
}
