package client

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gomessage/apps/models"
	"gomessage/apps/models/clients"
	"gomessage/apps/views"
	"net/http"
)

// PostClient
// @Tags Client
// @Summary 新增一个客户端
// @Router /api/v1/:namespace/client [POST]
func PostClient(g *gin.Context) {
	ns := g.Param("namespace")
	body := RequestBody{}
	if err := g.ShouldBindJSON(&body); err != nil {
		return
	}
	body.Namespace = ns

	switch body.ClientType {
	case "dingtalk":
		tmpClientData := RequestDataDingtalk{}
		if err := json.Unmarshal(body.ClientInfo, &tmpClientData); err != nil {
			return
		}
		var urls []string
		for _, v := range tmpClientData.RobotUrlList {
			urls = append(urls, v.Url)
		}
		tmpClientData.Dingtalk.RobotUrlsList = urls
		body.Client.ExtendDingtalk = tmpClientData.Dingtalk

	case "wechat":
		clientInfo := clients.WechatApplication{}
		if err := json.Unmarshal(body.ClientInfo, &clientInfo); err != nil {
			return
		}
		body.Client.ExtendWechatApplication = &clientInfo

	case "wechat_robot":
		clientInfo := RequestDataWechatRobot{}
		if err := json.Unmarshal(body.ClientInfo, &clientInfo); err != nil {
			return
		}
		var urls []string
		for _, v := range clientInfo.RobotUrlList {
			urls = append(urls, v.Url)
		}
		clientInfo.WechatRobot.RobotUrlsList = urls
		body.Client.ExtendWechatRobot = clientInfo.WechatRobot

	case "feishu":
		clientInfo := RequestDataFeishu{}
		if err := json.Unmarshal(body.ClientInfo, &clientInfo); err != nil {
			return
		}
		var urls []string
		for _, v := range clientInfo.RobotUrlList {
			urls = append(urls, v.Url)
		}
		clientInfo.Feishu.RobotUrlsList = urls
		body.Client.ExtendFeishu = clientInfo.Feishu

	default:
		return
	}

	result, err := models.AddClient(body.Client)
	if err != nil {
		return
	}

	g.JSON(http.StatusOK, views.ResponseSuccessful("创建成功", result))
}
