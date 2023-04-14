package client

import (
	"github.com/gin-gonic/gin"
	"gomessage/apps/models"
	"gomessage/apps/views"
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
		g.JSON(http.StatusBadRequest, views.ResponseFailure("查询错误", err))
	} else {
		respData := ResponseData{Client: client}
		if client.ClientType == "dingtalk" {
			var urls []OneUrl
			for _, urlAddress := range client.ExtendDingtalk.RobotUrlInfoList { //这里的RobotUrlInfoList，是从数据库取出的压缩数据，展开后得到的内容
				urls = append(urls, OneUrl{Url: urlAddress})
			}
			cInfo := ClientInfoDingtalk{
				Dingtalk:     client.ExtendDingtalk,
				RobotUrlList: urls,
			}
			respData.ClientInfo = cInfo

		} else if client.ClientType == "wechat" {
			client.ExtendWechatApplication.Secret = client.ExtendWechatApplication.Secret[:5] + "*****"
			respData.ClientInfo = client.ExtendWechatApplication

		} else if client.ClientType == "feishu" {
			var urls []OneUrl
			for _, v := range client.ExtendFeishu.RobotUrls {
				urls = append(urls, OneUrl{Url: v})
			}
			cInfo := ClientInfoFeishu{
				Feishu:       client.ExtendFeishu,
				RobotUrlList: urls,
			}
			respData.ClientInfo = cInfo
		}

		g.JSON(http.StatusOK, views.ResponseSuccessful("查询成功", respData))
	}
	return
}
