package client

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gomessage/models"
	"gomessage/pkg/general"
	"net/http"
	"strconv"
)

// PutClientActive
// @Tags Client
// @Summary 修改一个客户端
// @Router /api/v1/:namespace/client/:id [PUT]
func PutClientActive(g *gin.Context) {
	//更新客户端的激活状态
	id, _ := strconv.Atoi(g.Param("id"))
	client := models.Client{}
	err := g.ShouldBindJSON(&client)
	if err != nil {
		return
	}

	result, err := models.UpdateClientActive(id, &client)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	} else {
		g.JSON(http.StatusOK, result)
	}
}

// PutClientInfo
// @Tags Client
// @Summary 修改一个客户端
// @Router /api/v1/:namespace/client-info/:id [PUT]
func PutClientInfo(g *gin.Context) {

	//获取url中的客户端id
	id, _ := strconv.Atoi(g.Param("id"))

	//获取clientBody信息
	clientRequestBody := models.Client{}
	if err := g.ShouldBindJSON(&clientRequestBody); err != nil {
		return
	}

	//获取url中的namespace
	clientRequestBody.Namespace = g.Param("namespace")

	//判断客户端类型（主要是用来处理客客户端多url的问题）
	switch clientRequestBody.ClientType {
	case "dingtalk":
		//绑定钉钉客户端的信息
		if err := json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendDingtalk); err != nil {
			return
		}

	case "feishu":
		//绑定飞书客户端信息
		if err := json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendFeishu); err != nil {
			return
		}

	case "wechat_robot":
		//绑定微信机器人客户端信息
		wechatRobot := models.WechatRobot{}
		if err := json.Unmarshal(clientRequestBody.ClientInfo, &wechatRobot); err != nil {
			return
		}
		clientRequestBody.ExtendWechatRobot = &wechatRobot

	case "wechat":
		//微信应用号
		wechatApplication := models.WechatApplication{}
		if err := json.Unmarshal(clientRequestBody.ClientInfo, &wechatApplication); err != nil {
			return
		}
		clientRequestBody.ExtendWechatApplication = &wechatApplication

	default:
		return
	}

	result, err := models.UpdateClientInfo(id, &clientRequestBody)
	if err != nil {
		return
	}

	g.JSON(http.StatusOK, general.ResponseSuccessful("更新成功", result))
}
