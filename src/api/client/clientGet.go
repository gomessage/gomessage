package client

import (
	"github.com/gin-gonic/gin"
	models2 "gomessage/src/models"
	utils2 "gomessage/src/pkg/utils"
	"net/http"
	"strconv"
)

type RequestDataDingtalk struct {
	*models2.Dingtalk
	RobotUrlList []models2.Url `json:"robot_url_list"`
}

type RequestDataFeishu struct {
	*models2.Feishu
	RobotUrlList []models2.Url `json:"robot_url_list"`
}

type RequestDataWechatRobot struct {
	*models2.WechatRobot
	RobotUrlList []models2.Url `json:"robot_url_list"`
}

// GetClient
// @Tags Client
// @Summary 查询一个客户端
// @Router /api/v1/:namespace/client/:id [GET]
func GetClient(g *gin.Context) {
	type ResponseData struct {
		*models2.Client
		ClientInfo any `json:"client_info"`
	}

	id, _ := strconv.Atoi(g.Param("id"))
	client, err := models2.GetClientById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils2.ResponseFailure("查询错误", err))
	} else {
		respData := ResponseData{Client: client}

		switch client.ClientType {
		case utils2.VarDingtalk:
			var urls []models2.Url
			for _, urlAddress := range client.ExtendDingtalk.RobotUrlRandomList { //这里的RobotUrlInfoList，是从数据库取出的压缩数据，展开后得到的内容
				urls = append(urls, models2.Url{Url: urlAddress})
			}
			cInfo := RequestDataDingtalk{
				Dingtalk:     client.ExtendDingtalk,
				RobotUrlList: urls,
			}
			respData.ClientInfo = cInfo

		case utils2.VarFeishu:
			var urls []models2.Url
			for _, v := range client.ExtendFeishu.RobotUrlRandomList {
				urls = append(urls, models2.Url{Url: v})
			}
			cInfo := RequestDataFeishu{
				Feishu:       client.ExtendFeishu,
				RobotUrlList: urls,
			}
			respData.ClientInfo = cInfo

		case utils2.VarWechatRobot:
			var urls []models2.Url
			for _, v := range client.ExtendWechatRobot.RobotUrlRandomList {
				urls = append(urls, models2.Url{Url: v})
			}
			cInfo := RequestDataWechatRobot{
				WechatRobot:  client.ExtendWechatRobot,
				RobotUrlList: urls,
			}
			respData.ClientInfo = cInfo

		case utils2.VarWechatApplication:
			client.ExtendWechatApplication.Secret = client.ExtendWechatApplication.Secret[:5] + "*****"
			respData.ClientInfo = client.ExtendWechatApplication
		}

		g.JSON(http.StatusOK, utils2.ResponseSuccessful("查询成功", respData))
	}
	return
}
