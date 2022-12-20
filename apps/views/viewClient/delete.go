package viewClient

import (
	"github.com/gin-gonic/gin"
	"gomessage/apps/models"
	"net/http"
	"strconv"
)

// DeleteClient
// @Tags Client
// @Summary 删除一个客户端
// @Router /api/v1/:namespace/client/:id [DELETE]
func DeleteClient(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	_, err := models.DeleteClient(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	} else {
		g.JSON(http.StatusOK, map[string]bool{"result": true})
	}
}
