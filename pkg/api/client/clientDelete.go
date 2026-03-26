package client

import (
	"errors"
	"fmt"
	"gomessage/pkg/models"
	"gomessage/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteClient
// @Tags Client
// @Summary 删除一个客户端
// @Router /api/v1/:namespace/client/:id [DELETE]
func DeleteClient(g *gin.Context) {
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
	num, err := models.DeleteClient(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("删除失败", err))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("删除成功", fmt.Sprintf("受影响的行数为：%v", num)))
}
