package views

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gomessage/apps/controllers/core/v1"
	"gomessage/apps/controllers/core/v3"
	"gomessage/apps/controllers/hijacking"
	"gomessage/apps/models"
	"gomessage/utils/log/loggers"
	"io"
	"net/http"
	"time"
)

// GoMessageByTransport
// @Tags gomessage
// @Router /go/:namespace [POST]
func GoMessageByTransport(g *gin.Context) {

	/*
	 * TODO: 获取通道信息
	 */
	nsInfo := v1.GetNsInfo(g.Param("namespace"))

	/*
	 * TODO: 获取过境数据
	 */
	//获取请求数据
	hijacking.CacheData.RequestTime = time.Now()
	hijacking.CacheData.RequestByte, _ = io.ReadAll(g.Request.Body)                 //g.Request.Body中的数据只能读取一次，是因为"流"的指针被移位了
	g.Request.Body = io.NopCloser(bytes.NewBuffer(hijacking.CacheData.RequestByte)) //向g.Request.Body回写数据
	v3.FlatteningJson(hijacking.CacheData.RequestByte)                              //解析json
	if err := g.ShouldBindJSON(&hijacking.CacheData.RequestJson); err != nil {      //把请求数据绑定到CacheData.RequestJson
		return
	}

	/*
	 * TODO: 写入缓存便于劫持层读取信息
	 */
	hijacking.SetCacheData(nsInfo.Name, hijacking.CacheData) //把推送过来的数据写入缓存（一个命名空间中，同一时间只能写入一条数据）

	/*
	 * TODO: 获取通道的用户配置信息
	 */
	nsUserConfig := v1.GetNamespaceUserConfig(nsInfo.Name, nsInfo.IsRenders) //"通道自身信息"与"通道中用户添加的信息"不要搞混了

	/*
	 * TODO: 根据不同的客户端来产生对应的操作
	 */
	for _, client := range nsUserConfig.ActiveClient {
		var messages []any
		clientInfo, _ := models.GetClientById(client.ID)

		//获取interface的实例对象（该接口有两个方法：消息体处理的封装方法、推送消息的封装方法）
		var clientAction v3.ClientAction
		switch clientInfo.ClientType {
		case "dingtalk":
			clientAction = &v3.ClientActionDingtalk{Client: clientInfo}

		case "feishu":
			clientAction = &v3.ClientActionFeishu{Client: clientInfo}

		case "wechat":
			clientAction = &v3.ClientActionWechatApplication{
				CorpId:      clientInfo.ExtendWechatApplication.CorpId,
				AgentId:     clientInfo.ExtendWechatApplication.AgentId,
				AgentSecret: clientInfo.ExtendWechatApplication.Secret,
				Touser:      clientInfo.ExtendWechatApplication.Touser,
			}

		case "wechat_robot":
			clientAction = &v3.ClientActionWechatRobot{Client: clientInfo}

		default:
			g.JSON(http.StatusBadRequest, "客户端类型错误")
			return
		}

		/*
		 * TODO: 判断是否启用GoMessage的渲染层
		 * 这里的概念有点类似于网络七层模型中分层的概念：
		 * 转发层：类似于第四层的TCP/IP传输层
		 * 渲染层：类似于第七层的HTTP传输层
		 * 平时，有的数据过境GoMessage时都会100%的经过"转发层"，
		 * 如果此时用户打开"渲染开关"，则过境数据会被"渲染层"捞走，借用GoMessage强大的渲染能力把原始json信息转染成"人类可读"的很美观的信息格式
		 * 最后再下沉到"转发层"继续向后推送
		 */
		if nsInfo.IsRenders {
			//渲染出需要的"内容体"
			contentList := v3.RendersContentData(hijacking.CacheData.RequestByte, nsUserConfig.VariablesMap, nsUserConfig.MessageTemplate)
			//渲染出需要的"消息体"
			messages = clientAction.RendersMessages(clientInfo, nsUserConfig.IsMerge, contentList)
		} else {
			messages = append(messages, hijacking.CacheData.RequestJson)
		}

		// TODO：推送消息
		clientAction.PushMessages(messages) //入参只有一个"[]any"类型的值，后期可以基于该接口实现各种不同的"接收客户端"

		// TODO: 记录器（待实现...）
	}

	g.JSON(http.StatusOK, "ok")
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
