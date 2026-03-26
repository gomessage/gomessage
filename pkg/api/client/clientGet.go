package client

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gomessage/pkg/models"
	"gomessage/pkg/utils"
	"net/http"
	"strconv"
)

type RequestDataDingtalk struct {
	*models.Dingtalk
	RobotUrlList []models.Url `json:"robot_url_list"`
}

type RequestDataFeishu struct {
	*models.Feishu
	RobotUrlList []models.Url `json:"robot_url_list"`
}

type RequestDataWechatRobot struct {
	*models.WechatRobot
	RobotUrlList []models.Url `json:"robot_url_list"`
}

// GetClient
// @Tags Client
// @Summary 查询一个客户端
// @Router /api/v1/:namespace/client/:id [GET]
func GetClient(g *gin.Context) {
	type ResponseData struct {
		*models.Client
		ClientInfo any `json:"client_info"`
	}

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
		return
	}
	client, err := models.GetClientById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("查询错误", err))
		return
	}
	if client.Namespace != g.Param("namespace") {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", errors.New("客户端不属于当前命名空间")))
		return
	}
	respData := ResponseData{Client: client}

	switch client.ClientType {
	case utils.VarDingtalk:
		var urls []models.Url
		for _, urlAddress := range client.ExtendDingtalk.RobotUrlRandomList {
			urls = append(urls, models.Url{Url: urlAddress})
		}
		cInfo := RequestDataDingtalk{
			Dingtalk:     client.ExtendDingtalk,
			RobotUrlList: urls,
		}
		respData.ClientInfo = cInfo

	case utils.VarFeishu:
		var urls []models.Url
		for _, v := range client.ExtendFeishu.RobotUrlRandomList {
			urls = append(urls, models.Url{Url: v})
		}
		cInfo := RequestDataFeishu{
			Feishu:       client.ExtendFeishu,
			RobotUrlList: urls,
		}
		respData.ClientInfo = cInfo

	case utils.VarWechatRobot:
		var urls []models.Url
		for _, v := range client.ExtendWechatRobot.RobotUrlRandomList {
			urls = append(urls, models.Url{Url: v})
		}
		cInfo := RequestDataWechatRobot{
			WechatRobot:  client.ExtendWechatRobot,
			RobotUrlList: urls,
		}
		respData.ClientInfo = cInfo

	case utils.VarWechatApplication:
		respData.ClientInfo = client.ExtendWechatApplication
	}

	g.JSON(http.StatusOK, utils.ResponseSuccessful("查询成功", respData))
}
