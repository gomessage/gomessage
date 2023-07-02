package client

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomessage/models"
	"gomessage/pkg/utils"
	"net/http"
	"strconv"
)

// DeleteClient
// @Tags Client
// @Summary 删除一个客户端
// @Router /api/v1/:namespace/client/:id [DELETE]
func DeleteClient(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	num, err := models.DeleteClient(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("删除失败", err))
	} else {
		g.JSON(http.StatusOK, utils.ResponseSuccessful("删除成功", fmt.Sprintf("受影响的行数为：%v", num)))
	}
}
