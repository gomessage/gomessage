package views

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gomessage/apps/controllers/hijacking"
	"gomessage/apps/controllers/send"
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

	/*
	 *
	 * TODO: 获取通道信息
	 *
	 */
	namespaceInfo := send.GetNs(g.Param("namespace"))

	/*
	 *
	 * TODO: 获取过境数据
	 *
	 */
	//获取请求数据
	hijacking.CacheData.RequestTime = time.Now()
	hijacking.CacheData.RequestByte, _ = io.ReadAll(g.Request.Body)                 //g.Request.Body中的数据只能读取一次，是因为"流"的指针被移位了
	g.Request.Body = io.NopCloser(bytes.NewBuffer(hijacking.CacheData.RequestByte)) //向g.Request.Body回写数据
	if err := g.ShouldBindJSON(&hijacking.CacheData.RequestJson); err != nil {      //把请求数据绑定到CacheData.RequestJson
		return
	}

	/*
	 *
	 * TODO: 写入缓存便于劫持层读取信息
	 *
	 */
	//把推送过来的数据写入缓存（一个命名空间中，同一时间只能写入一条数据）
	hijacking.SetCacheData(namespaceInfo.Name, hijacking.CacheData)

	/*
	 *
	 * TODO: 获取通道的用户配置信息
	 *
	 */
	//从数据库中拿到用户当前用户在图形界面上配置的参数
	nsUserConfig := send.GetNamespaceUserConfig(namespaceInfo.Name)

	/*
	 *
	 * TODO: 渲染数据
	 *
	 */
	rendersDataList := send.RendersRequestData(namespaceInfo.IsRenders, nsUserConfig, hijacking.CacheData.RequestByte)

	/*
	 *
	 * TODO: 对应客户端的消息体组装
	 *
	 */
	readyClientList := send.AssembledMessage(namespaceInfo.IsRenders, nsUserConfig, rendersDataList)

	/*
	 *
	 * TODO: 推送数据
	 *
	 */
	for _, readyClient := range readyClientList {
		url := readyClient.Url
		data := readyClient.Data
		for _, d := range data {
			send.Push(d, url)
		}
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
