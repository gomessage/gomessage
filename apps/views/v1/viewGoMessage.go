package v1

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gomessage/apps/controllers/core/interfaces"
	"gomessage/apps/controllers/core/realized/v1"
	"gomessage/apps/controllers/hijacking"
	"gomessage/apps/controllers/send"
	"gomessage/apps/models"
	"io"
	"net/http"
	"time"
)

// GoMessageByPost
// @Tags gomessage
// @Router /go/:namespace [POST]
func GoMessageByPost(g *gin.Context) {

	/*
	 *
	 * TODO: 获取通道信息
	 *
	 */
	namespaceInfo := send.GetNs(g.Param("namespace"))

	/*
	 *
	 * TODO: 获取过境数据
	 *
	 */
	//获取请求数据
	hijacking.CacheData.RequestTime = time.Now()
	hijacking.CacheData.RequestByte, _ = io.ReadAll(g.Request.Body)                 //g.Request.Body中的数据只能读取一次，是因为"流"的指针被移位了
	g.Request.Body = io.NopCloser(bytes.NewBuffer(hijacking.CacheData.RequestByte)) //向g.Request.Body回写数据
	if err := g.ShouldBindJSON(&hijacking.CacheData.RequestJson); err != nil {      //把请求数据绑定到CacheData.RequestJson
		return
	}

	/*
	 *
	 * TODO: 写入缓存便于劫持层读取信息
	 *
	 */
	//把推送过来的数据写入缓存（一个命名空间中，同一时间只能写入一条数据）
	hijacking.SetCacheData(namespaceInfo.Name, hijacking.CacheData)

	/*
	 *
	 * TODO: 获取通道的用户配置信息
	 *
	 */
	//从数据库中拿到用户当前用户在图形界面上配置的参数
	namespaceUserConfig := send.GetNamespaceUserConfig(namespaceInfo.Name)

	/*
	 *
	 * TODO: 渲染数据
	 *
	 */
	//渲染出来的"内容体"与客户端类型无关，只渲染一次，所有类型的客户端都可以使用
	rendersData := &v1.GetRendersResult{Rds: namespaceInfo.IsRenders}

	/*
	 *
	 * TODO: 处理发送行为
	 *
	 */
	//rendersDataList := send.RendersRequestData(namespaceInfo.GetRendersResult, namespaceUserConfig, hijacking.CacheData.RequestByte)
	var action *interfaces.Action

	//遍历当前通道下已经被激活的客户端（从这一步开始，后续的所有行为，要根据不同的客户端来动态实例化接口）
	for _, client := range namespaceUserConfig.ActiveClient {
		//获取客户端详情
		clientInfo, _ := models.GetClientById(client.ID)

		//根据客户端类型加载不同的实例化对象
		switch clientInfo.ClientType {

		case "dingtalk":
			action = interfaces.NewAction(
				rendersData,
				&v1.DingtalkMessageAssembled{},
				&v1.GeneralPush{},
				&v1.GeneralRecord{},
			)

		case "feishu":
			action = interfaces.NewAction(
				rendersData,
				&v1.FeishuMessageAssembled{},
				&v1.GeneralPush{},
				&v1.GeneralRecord{},
			)

		case "wechat":
			action = interfaces.NewAction(
				rendersData,
				&v1.WechatMessageAssembled{},
				&v1.WechatPush{
					CorpId:      clientInfo.ExtendWechat.CorpId,
					AgentId:     clientInfo.ExtendWechat.AgentId,
					AgentSecret: clientInfo.ExtendWechat.Secret,
					Touser:      clientInfo.ExtendWechat.Touser,
				},
				&v1.GeneralRecord{})

		default:
			action = interfaces.NewAction(
				rendersData,
				&v1.GeneralMessageAssembled{},
				&v1.GeneralPush{},
				&v1.GeneralRecord{},
			)

		}
		err := action.Working(namespaceInfo.IsRenders, hijacking.CacheData.RequestByte, namespaceUserConfig, clientInfo)
		if err != nil {
			g.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	g.JSON(http.StatusOK, "ok")
}
