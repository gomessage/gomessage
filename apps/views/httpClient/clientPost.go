package httpClient

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gomessage/apps/models"
	"gomessage/apps/views/httpBase"
	"net/http"
)

// PostClient
// @Tags Client
// @Summary 新增一个客户端
// @Router /api/v1/:namespace/client [POST]
func PostClient(g *gin.Context) {
	ns := g.Param("namespace")
	body := RequestBody{}
	if err := g.ShouldBindJSON(&body); err != nil {
		return
	}
	body.Namespace = ns

	switch body.ClientType {
	/*
	 * 解析钉钉客户端ClientInfo中的数据
	 */
	case "dingtalk":
		tmpClientData := ClientInfoDingtalk{}
		if err := json.Unmarshal(body.ClientInfo, &tmpClientData); err != nil {
			return
		}
		var urls []string
		for _, v := range tmpClientData.RobotUrlList {
			urls = append(urls, v.Url)
		}
		tmpClientData.Dingtalk.RobotUrlInfoList = urls
		body.Client.ExtendDingtalk = tmpClientData.Dingtalk

	/*
	 * 解析微信客户端ClientInfo中的数据
	 */
	case "wechat":
		clientInfo := models.Wechat{}
		if err := json.Unmarshal(body.ClientInfo, &clientInfo); err != nil {
			return
		}
		body.Client.ExtendWechat = &clientInfo

	/*
	 * 解析飞书客户端ClientInfo中的数据
	 */
	case "feishu":
		clientInfo := ClientInfoFeishu{}
		if err := json.Unmarshal(body.ClientInfo, &clientInfo); err != nil {
			return
		}
		var urls []string
		for _, v := range clientInfo.RobotUrlList {
			urls = append(urls, v.Url)
		}
		clientInfo.Feishu.RobotUrls = urls
		body.Client.ExtendFeishu = clientInfo.Feishu

	default:
		return
	}

	result, err := models.AddClient(body.Client)
	if err != nil {
		return
	}

	//g.JSON(http.StatusOK, ResponseData{
	//    Result: true,
	//    Client: result,
	//})

	g.JSON(http.StatusOK, httpBase.ResponseSuccessful("创建成功", result))
}
