package client

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gomessage/pkg/models"
	"gomessage/pkg/utils"
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
	case utils.VarDingtalk:
		json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendDingtalk)

	case utils.VarWechatRobot:
		json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendWechatRobot)

	case utils.VarFeishu:
		json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendFeishu)

	case utils.VarWechatApplication:
		json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendWechatApplication)

	default:
		return
	}

	result, err := models.AddClient(&clientRequestBody)
	if err != nil {
		return
	}

	g.JSON(http.StatusOK, utils.ResponseSuccessful("创建成功", result))
}
