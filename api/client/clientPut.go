package client

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gomessage/models"
	"gomessage/pkg/utils"
	"net/http"
	"strconv"
)

// PutClientInfo
// @Tags Client
// @Summary 修改一个客户端
// @Router /api/v1/:namespace/client-info/:id [PUT]
func PutClientInfo(g *gin.Context) {
	clientRequestBody := models.Client{}
	if err := g.ShouldBindJSON(&clientRequestBody); err != nil {
		return
	}
	clientRequestBody.Namespace = g.Param("namespace")
	switch clientRequestBody.ClientType {
	case utils.VarDingtalk:
		json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendDingtalk)

	case utils.VarFeishu:
		json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendFeishu)

	case utils.VarWechatRobot:
		json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendWechatRobot)

	case utils.VarWechatApplication:
		json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendWechatApplication)

	default:
		return
	}

	id, _ := strconv.Atoi(g.Param("id"))
	result, err := models.UpdateClientInfo(id, &clientRequestBody)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("更新失败", err))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("更新成功", result))
}

// PutClientActive
// @Tags Client
// @Summary 修改客户端激活状态
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
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("", err))
	} else {
		g.JSON(http.StatusOK, utils.ResponseSuccessful("", result))
	}
}
