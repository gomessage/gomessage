package api

import (
	"bytes"
	"fmt"
	"gomessage/pkg/models"
	v1 "gomessage/pkg/services/core/v1"
	v3 "gomessage/pkg/services/core/v3"
	"gomessage/pkg/services/hijacking"
	"gomessage/pkg/utils"
	"gomessage/pkg/utils/log/loggers"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GoMessageByTransport
// @Tags gomessage
// @Router /go/:namespace [POST]
func GoMessageByTransport(g *gin.Context) {
	//TODO：劫持数据的容器
	hijackingData := hijacking.ArbitrarilyJsonData{}

	//TODO: 获取通道对象
	nsObj := v1.GetNsInfo(g.Param("namespace"))

	//TODO: 获取过境数据
	hijackingData.RequestTime = time.Now()
	requestBytes, err := io.ReadAll(g.Request.Body)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("读取请求体失败", err))
		return
	}
	hijackingData.RequestByte = requestBytes
	g.Request.Body = io.NopCloser(bytes.NewBuffer(hijackingData.RequestByte)) //向g.Request.Body回写数据
	if err := g.ShouldBindJSON(&hijackingData.RequestJson); err != nil {      //把请求数据绑定到CacheData.RequestJson
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求体必须是合法JSON", err))
		return
	}

	//TODO：扁平化解析json
	//v3.FlatteningJson(hijackingData.RequestByte)                              //扁平化解析json

	loggers.PushLogger.WithFields(logrus.Fields{
		"namespace": nsObj.Name,
		"content":   string(hijackingData.RequestByte),
	}).Info("上游消息")

	//TODO: 写入缓存便于劫持层读取信息
	hijacking.SetCacheData(nsObj.Name, hijackingData) //把推送过来的数据写入缓存（一个命名空间中，同一时间只能写入一条数据）

	//TODO: 获取通道的用户配置信息
	nsUserConfig, err := v1.GetNamespaceUserConfig(nsObj.Name, nsObj.IsRenders) //"通道自身信息"与"通道中用户添加的信息"不要搞混了
	if err != nil {
		g.JSON(http.StatusInternalServerError, utils.ResponseFailure("通道配置获取失败", err))
		return
	}

	successCount := 0
	failureCount := 0
	var failureReasons []string

	//TODO: 根据不同的客户端来产生对应的操作
	for _, client := range nsUserConfig.ActiveClient {
		var messages []any
		clientInfo, err := models.GetClientById(client.ID)
		if err != nil {
			loggers.DefaultLogger.Errorln("查询客户端失败：", err)
			failureCount++
			failureReasons = append(failureReasons, fmt.Sprintf("client_id=%d 查询失败", client.ID))
			continue
		}

		//获取interface的实例对象（该接口有两个方法：消息体处理的封装方法、推送消息的封装方法）
		var clientAction v3.ClientAction
		switch clientInfo.ClientType {
		case utils.VarDingtalk:
			clientAction = &v3.ClientActionDingtalk{Client: clientInfo}

		case utils.VarFeishu:
			clientAction = &v3.ClientActionFeishu{Client: clientInfo}

		case utils.VarWechatApplication:
			if clientInfo.ExtendWechatApplication == nil {
				loggers.DefaultLogger.Errorln("企业微信客户端扩展信息为空，client_id=", clientInfo.ID)
				continue
			}
			clientAction = &v3.ClientActionWechatApplication{
				CorpId:      clientInfo.ExtendWechatApplication.CorpId,
				AgentId:     clientInfo.ExtendWechatApplication.AgentId,
				AgentSecret: clientInfo.ExtendWechatApplication.Secret,
				Touser:      clientInfo.ExtendWechatApplication.Touser,
			}

		case utils.VarWechatRobot:
			clientAction = &v3.ClientActionWechatRobot{Client: clientInfo}

		default:
			loggers.DefaultLogger.Errorln("客户端类型错误：", clientInfo.ClientType)
			failureCount++
			failureReasons = append(failureReasons, fmt.Sprintf("client_id=%d 客户端类型错误", clientInfo.ID))
			continue
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
		if nsObj.IsRenders {
			//渲染出需要的"内容体"
			contentList := v3.RendersContentData(hijackingData.RequestByte, nsUserConfig.VariablesMap, nsUserConfig.MessageTemplate)
			if len(contentList) == 0 {
				failureCount++
				failureReasons = append(failureReasons, fmt.Sprintf("client_id=%d 渲染结果为空", clientInfo.ID))
				continue
			}
			//渲染出需要的"消息体"
			messages = clientAction.RendersMessages(clientInfo, nsUserConfig.IsMerge, contentList)
		} else {
			messages = append(messages, hijackingData.RequestJson)
		}
		if len(messages) == 0 {
			failureCount++
			failureReasons = append(failureReasons, fmt.Sprintf("client_id=%d 消息体为空", clientInfo.ID))
			continue
		}

		// TODO：推送消息
		s, f := clientAction.PushMessages(messages) //入参只有一个"[]any"类型的值，后期可以基于该接口实现各种不同的"接收客户端"
		successCount += s
		failureCount += f

		// TODO: 记录器（待实现...）
	}
	if successCount == 0 {
		reason := strings.Join(failureReasons, "; ")
		g.JSON(http.StatusInternalServerError, utils.ResponseFailure(fmt.Sprintf("推送失败：成功%d，失败%d", successCount, failureCount), fmt.Errorf(reason)))
		return
	}
	if failureCount > 0 {
		g.JSON(http.StatusOK, utils.ResponseSuccessful(fmt.Sprintf("部分成功：成功%d，失败%d", successCount, failureCount), "ok"))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("推送完成", "ok"))
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
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("namespace不存在", err))
	} else {
		g.JSON(http.StatusOK, utils.ResponseSuccessful("namespace ready", result))
	}
}
