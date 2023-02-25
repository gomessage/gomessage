package httpClient

import (
    "github.com/gin-gonic/gin"
    "gomessage/apps/models"
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
        g.JSON(http.StatusBadRequest, "参数错误")
    } else {
        respData := ResponseData{Client: client}
        if client.ClientType == "dingtalk" {
            var urls []Url
            for _, v := range client.ClientInfoDingtalk.RobotUrls {
                urls = append(urls, Url{Url: v})
            }
            cInfo := ClientInfoDingtalk{
                Dingtalk:     client.ClientInfoDingtalk,
                RobotUrlList: urls,
            }
            respData.ClientInfo = cInfo

        } else if client.ClientType == "wechat" {
            client.ClientInfoWechat.Secret = client.ClientInfoWechat.Secret[:5] + "*****"
            respData.ClientInfo = client.ClientInfoWechat

        } else if client.ClientType == "feishu" {
            var urls []Url
            for _, v := range client.ClientInfoFeishu.RobotUrls {
                urls = append(urls, Url{Url: v})
            }
            cInfo := ClientInfoFeishu{
                Feishu:       client.ClientInfoFeishu,
                RobotUrlList: urls,
            }
            respData.ClientInfo = cInfo
        }

        g.JSON(http.StatusOK, respData)
    }
    return
}
