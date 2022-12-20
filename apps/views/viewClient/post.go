package viewClient

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gomessage/apps/models"
	"net/http"
)

// PostClient
// @Tags Client
// @Summary 新增一个客户端
// @Router /api/v1/:namespace/client [POST]
func PostClient(g *gin.Context) {
	type ResponseData struct {
		*models.Client
		Result bool `json:"result"`
	}

	ns := g.Param("namespace")
	request := Request{}
	if err := g.ShouldBindJSON(&request); err != nil {
		return
	}
	request.Namespace = ns

	switch request.ClientType {
	/*
	 * 解析钉钉客户端ClientInfo中的数据
	 */
	case "dingtalk":
		clientInfo := ClientInfoDingtalk{}
		if err := json.Unmarshal(request.ClientInfo, &clientInfo); err != nil {
			return
		}
		var urls []string
		for _, v := range clientInfo.RobotUrlList {
			urls = append(urls, v.Url)
		}
		clientInfo.Dingtalk.RobotUrls = urls
		request.Client.ClientInfoDingtalk = clientInfo.Dingtalk

	/*
	 * 解析微信客户端ClientInfo中的数据
	 */
	case "wechat":
		clientInfo := models.Wechat{}
		if err := json.Unmarshal(request.ClientInfo, &clientInfo); err != nil {
			return
		}
		request.Client.ClientInfoWechat = &clientInfo

	/*
	 * 解析飞书客户端ClientInfo中的数据
	 */
	case "feishu":
		clientInfo := ClientInfoFeishu{}
		if err := json.Unmarshal(request.ClientInfo, &clientInfo); err != nil {
			return
		}
		var urls []string
		for _, v := range clientInfo.RobotUrlList {
			urls = append(urls, v.Url)
		}
		clientInfo.Feishu.RobotUrls = urls
		request.Client.ClientInfoFeishu = clientInfo.Feishu

	default:
		return
	}

	result, err := models.AddClient(request.Client)
	if err != nil {
		return
	}

	g.JSON(http.StatusOK, ResponseData{
		Result: true,
		Client: result,
	})
}
