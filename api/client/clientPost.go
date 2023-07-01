package client

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gomessage/models"
	"gomessage/pkg/general"
	"net/http"
)

// PostClient
// @Tags Client
// @Summary 新增一个客户端
// @Router /api/v1/:namespace/client [POST]
func PostClient(g *gin.Context) {
	//获取clientBody信息
	clientRequestBody := models.Client{}
	if err := g.ShouldBindJSON(&clientRequestBody); err != nil {
		return
	}

	//获取url中的namespace
	clientRequestBody.Namespace = g.Param("namespace")

	//判断客户端类型（绑定客户端信息到client的Extend延伸信息中）
	switch clientRequestBody.ClientType {
	case "dingtalk":
		//钉钉机器人
		dingtalk := models.Dingtalk{}
		json.Unmarshal(clientRequestBody.ClientInfo, &dingtalk)
		clientRequestBody.ExtendDingtalk = &dingtalk

	case "wechat_robot":
		//微信机器人
		wechatRobot := models.WechatRobot{}
		json.Unmarshal(clientRequestBody.ClientInfo, &wechatRobot)
		clientRequestBody.ExtendWechatRobot = &wechatRobot

	case "feishu":
		//飞书机器人
		feishu := models.Feishu{}
		json.Unmarshal(clientRequestBody.ClientInfo, &feishu)
		clientRequestBody.ExtendFeishu = &feishu

	case "wechat":
		//微信应用号
		wechatApplication := models.WechatApplication{}
		json.Unmarshal(clientRequestBody.ClientInfo, &wechatApplication)
		clientRequestBody.ExtendWechatApplication = &wechatApplication

	default:
		return
	}

	result, err := models.AddClient(&clientRequestBody)
	if err != nil {
		return
	}

	g.JSON(http.StatusOK, general.ResponseSuccessful("创建成功", result))
}
