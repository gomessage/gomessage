package client

import (
	"github.com/gin-gonic/gin"
	"gomessage/models"
	"gomessage/pkg/general"
	"net/http"
	"strconv"
)

// GetClient
// @Tags Client
// @Summary 查询一个客户端
// @Router /api/v1/:namespace/client/:id [GET]
func GetClient(g *gin.Context) {
	type ResponseData struct {
		*models.Client
		ClientInfo any `json:"client_info"`
	}

	id, _ := strconv.Atoi(g.Param("id"))
	client, err := models.GetClientById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, general.ResponseFailure("查询错误", err))
	} else {
		respData := ResponseData{Client: client}

		switch client.ClientType {
		case "dingtalk":
			var urls []OneUrl
			for _, urlAddress := range client.ExtendDingtalk.RobotUrlInfoList { //这里的RobotUrlInfoList，是从数据库取出的压缩数据，展开后得到的内容
				urls = append(urls, OneUrl{Url: urlAddress})
			}
			cInfo := RequestDataDingtalk{
				Dingtalk:     client.ExtendDingtalk,
				RobotUrlList: urls,
			}
			respData.ClientInfo = cInfo

		case "feishu":
			var urls []OneUrl
			for _, v := range client.ExtendFeishu.RobotUrlInfoList {
				urls = append(urls, OneUrl{Url: v})
			}
			cInfo := RequestDataFeishu{
				Feishu:       client.ExtendFeishu,
				RobotUrlList: urls,
			}
			respData.ClientInfo = cInfo

		case "wechat_robot":
			var urls []OneUrl
			for _, v := range client.ExtendWechatRobot.RobotUrlInfoList {
				urls = append(urls, OneUrl{Url: v})
			}
			cInfo := RequestDataWechatRobot{
				WechatRobot:  client.ExtendWechatRobot,
				RobotUrlList: urls,
			}
			respData.ClientInfo = cInfo

		case "wechat":
			client.ExtendWechatApplication.Secret = client.ExtendWechatApplication.Secret[:5] + "*****"
			respData.ClientInfo = client.ExtendWechatApplication
		}

		g.JSON(http.StatusOK, general.ResponseSuccessful("查询成功", respData))
	}
	return
}
