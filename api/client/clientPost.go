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
		json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendDingtalk)

	case "wechat_robot":
		//微信机器人
		json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendWechatRobot)

	case "feishu":
		//飞书机器人
		json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendFeishu)

	case "wechat":
		//微信应用号
		json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendWechatApplication)

	default:
		return
	}

	result, err := models.AddClient(&clientRequestBody)
	if err != nil {
		return
	}

	g.JSON(http.StatusOK, general.ResponseSuccessful("创建成功", result))
}
