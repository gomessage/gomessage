package client

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"gomessage/pkg/models"
	"gomessage/pkg/utils"
	"net/http"
	"strconv"
)

// PutClientInfo
// @Tags Client
// @Summary 修改一个客户端
// @Router /api/v1/:namespace/client-info/:id [PUT]
func PutClientInfo(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
		return
	}
	exist, err := models.GetClientLiteById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("查询失败", err))
		return
	}
	if exist.Namespace != g.Param("namespace") {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", errors.New("客户端不属于当前命名空间")))
		return
	}

	clientRequestBody := models.Client{}
	if err := g.ShouldBindJSON(&clientRequestBody); err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}
	clientRequestBody.Namespace = g.Param("namespace")
	switch clientRequestBody.ClientType {
	case utils.VarDingtalk:
		if err = json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendDingtalk); err != nil {
			g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
			return
		}

	case utils.VarFeishu:
		if err = json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendFeishu); err != nil {
			g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
			return
		}

	case utils.VarWechatRobot:
		if err = json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendWechatRobot); err != nil {
			g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
			return
		}

	case utils.VarWechatApplication:
		if err = json.Unmarshal(clientRequestBody.ClientInfo, &clientRequestBody.ExtendWechatApplication); err != nil {
			g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
			return
		}

	default:
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("客户端类型错误", errors.New("未知的client_type")))
		return
	}

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
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
		return
	}
	exist, err := models.GetClientLiteById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("查询失败", err))
		return
	}
	if exist.Namespace != g.Param("namespace") {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", errors.New("客户端不属于当前命名空间")))
		return
	}

	client := models.Client{}
	if err = g.ShouldBindJSON(&client); err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}

	result, err := models.UpdateClientActive(id, &client)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("", err))
	} else {
		g.JSON(http.StatusOK, utils.ResponseSuccessful("", result))
	}
}
