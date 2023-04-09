package views

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gomessage/apps/controllers/hijacking"
	"gomessage/apps/controllers/sends"
	"gomessage/apps/models"
	"gomessage/utils/log/loggers"
	"io"
	"net/http"
	"time"
)

// GoMessageByPost
// @Tags gomessage
// @Router /go/:namespace [POST]
func GoMessageByPost(g *gin.Context) {
	namespace := g.Param("namespace")
	if namespace == "message" {
		namespace = "default"
	}
	loggers.DefaultLogger.Info("消息发送至" + namespace + "命名空间")

	nsObj, _ := models.GetNamespaceByName(namespace)

	//如果开启数据渲染，则走渲染模式
	if nsObj.IsRenders {
		//获取请求数据
		hijacking.CacheData.RequestTime = time.Now()
		hijacking.CacheData.RequestByte, _ = io.ReadAll(g.Request.Body)                 //g.Request.Body中的数据只能读取一次，是因为"流"的指针被移位了
		g.Request.Body = io.NopCloser(bytes.NewBuffer(hijacking.CacheData.RequestByte)) //向g.Request.Body回写数据

		//把请求数据绑定到CacheData.RequestJson
		if err := g.ShouldBindJSON(&hijacking.CacheData.RequestJson); err != nil {
			return
		}

		//把推送过来的数据写入缓存（一个命名空间中，同一时间只能写入一条数据）
		hijacking.SetCacheData(namespace, hijacking.CacheData)

		//从数据库中拿到用户当前用户在图形界面上配置的参数
		userConfig := sends.GetUserConfig(namespace)

		//创建过境数据与用户变量之间的映射
		analysisDataList := sends.AnalysisData(userConfig.VariablesMap, hijacking.CacheData.RequestByte)

		//得到渲染完成后的消息列表
		templateMessageList := sends.CompleteMessage(userConfig.MessageTemplate, analysisDataList)

		//判断消息是否聚合发送
		if userConfig.MessageMerge {
			sends.Merge(templateMessageList, userConfig)
		} else {
			sends.Disperse(templateMessageList, userConfig)
		}

		g.JSON(http.StatusOK, "ok")
	} else {
		//如果没有开启渲染开关，则走"转发模式"

	}
}

// GoMessageByGet
// @Tags gomessage
// @Router /go/message [GET]
func GoMessageByGet(g *gin.Context) {
	namespace := g.Param("namespace")
	if namespace == "message" {
		namespace = "default"
	}
	loggers.DefaultLogger.Info("当前命名空间为：", namespace)

	result, err := models.GetNamespaceByName(namespace)
	if err != nil {
		g.JSON(http.StatusBadRequest, ResponseFailure("namespace不存在", err))
	} else {
		g.JSON(http.StatusOK, ResponseSuccessful("namespace ready", result))
	}
}
