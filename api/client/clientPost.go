package client

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gomessage/api"
	"gomessage/models"
	"gomessage/models/clients"
	"net/http"
)

// PostClient
// @Tags Client
// @Summary 新增一个客户端
// @Router /api/v1/:namespace/client [POST]
func PostClient(g *gin.Context) {
	body := RequestBody{}
	if err := g.ShouldBindJSON(&body); err != nil {
		return
	}
	body.Namespace = g.Param("namespace")

	switch body.ClientType {

	case "dingtalk":
		requestClient := RequestDataDingtalk{}
		if err := json.Unmarshal(body.ClientInfo, &requestClient); err != nil {
			return
		}
		var urls []string
		for _, v := range requestClient.RobotUrlList {
			urls = append(urls, v.Url)
		}
		requestClient.Dingtalk.RobotUrlInfoList = urls
		body.Client.ExtendDingtalk = requestClient.Dingtalk

	case "wechat_robot":
		requestClient := RequestDataWechatRobot{}
		if err := json.Unmarshal(body.ClientInfo, &requestClient); err != nil {
			return
		}
		var urls []string
		for _, v := range requestClient.RobotUrlList {
			urls = append(urls, v.Url)
		}
		requestClient.WechatRobot.RobotUrlInfoList = make([]string, 0)
		requestClient.WechatRobot.RobotUrlInfoList = urls
		body.Client.ExtendWechatRobot = requestClient.WechatRobot

	case "feishu":
		requestClient := RequestDataFeishu{}
		if err := json.Unmarshal(body.ClientInfo, &requestClient); err != nil {
			return
		}
		var urls []string
		for _, v := range requestClient.RobotUrlList {
			urls = append(urls, v.Url)
		}
		requestClient.Feishu.RobotUrlInfoList = urls
		body.Client.ExtendFeishu = requestClient.Feishu

	case "wechat":
		requestClient := clients.WechatApplication{}
		if err := json.Unmarshal(body.ClientInfo, &requestClient); err != nil {
			return
		}
		body.Client.ExtendWechatApplication = &requestClient

	default:
		return
	}

	result, err := models.AddClient(body.Client)
	if err != nil {
		return
	}

	g.JSON(http.StatusOK, api.ResponseSuccessful("创建成功", result))
}
