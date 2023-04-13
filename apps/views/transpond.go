package views

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gomessage/apps/controllers/hijacking"
	"gomessage/apps/controllers/send"
	"gomessage/apps/controllers/utils/base"
	"gomessage/apps/models"
	"io"
	"net/http"
	"time"
)

// GoMessageByTranspond
// @Tags gomessage
// @Router /go/:namespace [POST]
func GoMessageByTranspond(g *gin.Context) {

	/*
	 * TODO: 获取通道信息
	 */
	namespaceInfo := send.GetNs(g.Param("namespace"))

	/*
	 * TODO: 获取过境数据
	 */
	//获取请求数据
	hijacking.CacheData.RequestTime = time.Now()
	hijacking.CacheData.RequestByte, _ = io.ReadAll(g.Request.Body)                 //g.Request.Body中的数据只能读取一次，是因为"流"的指针被移位了
	g.Request.Body = io.NopCloser(bytes.NewBuffer(hijacking.CacheData.RequestByte)) //向g.Request.Body回写数据
	if err := g.ShouldBindJSON(&hijacking.CacheData.RequestJson); err != nil {      //把请求数据绑定到CacheData.RequestJson
		return
	}

	/*
	 * TODO: 写入缓存便于劫持层读取信息
	 */
	//把推送过来的数据写入缓存（一个命名空间中，同一时间只能写入一条数据）
	hijacking.SetCacheData(namespaceInfo.Name, hijacking.CacheData)

	/*
	 * TODO: 获取通道的用户配置信息
	 */
	//从数据库中拿到用户当前用户在图形界面上配置的参数
	namespaceUserConfig := send.GetNamespaceUserConfig(namespaceInfo.Name)

	for _, client := range namespaceUserConfig.ActiveClient {
		clientInfo, _ := models.GetClientById(client.ID)

		//获取interface的实例化对象（应该有两个方法：消息体封装、推送封装）
		var clientAction base.ClientAction
		clientAction = &base.ClientActionDingtalk{}

		// TODO: 判断是否启用GoMessage的渲染层，如果是执行渲染层
		//如果需要渲染（入参是[]byte，出参应该就是下文的data）
		if namespaceInfo.IsRenders {
			//判断客户端类型，渲染成指定的"内容体"，直接产出一个data（应该是一个interface类型，应该只有一个实现）
			contentList := RendersContentData(hijacking.CacheData.RequestByte, namespaceUserConfig.VariablesMap, namespaceUserConfig.MsgTemplate)

			//判断客户端类型，把"内容体"渲染成"消息体"（也应该是一个interface类型，应该又几个客户端，就有几个实现）
			clientAction.RendersMessages()
		}

		// TODO：调用推送消息的方法（通用、或微信专用）
		// 推送这里是一个interface类型（入参：有且只有一个data，这个地方只能是一个"消息体"）
		clientAction.PushMessages()

		// TODO: 记录器
		// 记录一条发送概况，用来展示统计信息
	}

	g.JSON(http.StatusOK, "ok")
}

func RendersContentData(requestByte []byte, varMap []map[string]string, msgTemplate string) []string {
	analysisDataList := send.AnalysisData(varMap, requestByte) //得到变量映射
	return send.CompleteMessage(msgTemplate, analysisDataList) //得到内容体
}
