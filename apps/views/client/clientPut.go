package client

import (
	"github.com/gin-gonic/gin"
	"gomessage/apps/models"
	"net/http"
	"strconv"
)

// PutClient
// @Tags Client
// @Summary 修改一个客户端
// @Router /api/v1/:namespace/client/:id [PUT]
func PutClient(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	client := models.Client{}
	err := g.ShouldBindJSON(&client)
	if err != nil {
		return
	}

	result, err := models.UpdateClientActive(id, &client)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	} else {
		g.JSON(http.StatusOK, result)
	}
}
