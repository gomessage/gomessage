package views

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomessage/apps/models"
	"net/http"
	"strconv"
)

// ListNamespace
// @Tags Namespace
// @Summary 获取所有命名空间
// @Router /api/v1/namespace [GET]
func ListNamespace(g *gin.Context) {
	isActive := g.DefaultQuery("is_active", "")
	switch isActive {
	case "true", "false", "1", "0", "":
		list, err := models.ListNamespace(isActive)
		if err != nil {
			g.JSON(http.StatusInternalServerError, "服务器内部错误")
		}
		g.JSON(http.StatusOK, list)

	default:
		g.JSON(http.StatusBadRequest, "参数错误")
	}
}

// PostNamespace
// @Tags Namespace
// @Summary 新增一个命名空间
// @Router /api/v1/namespace [POST]
func PostNamespace(g *gin.Context) {
	body := models.Namespace{}
	if err := g.ShouldBindJSON(&body); err != nil {
		return
	}
	namespace, err := models.AddNamespace(&body)
	if err != nil {
		g.JSON(http.StatusBadRequest, ResponseFailure("命名空间已存在，不能重复创建", err))
	} else {
		g.JSON(http.StatusOK, ResponseSuccessful("命名空间创建成功", &namespace))
	}
}

// GetNamespace
// @Tags Namespace
// @Summary 查询一个命名空间
// @Router /api/v1/namespace/:id [GET]
func GetNamespace(g *gin.Context) {
	//id, _ := strconv.ParseInt(g.Param("id"), 10, 64)
	id, _ := strconv.Atoi(g.Param("id"))
	result, err := models.GetNamespaceById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, "参数错误")
	} else {
		g.JSON(http.StatusOK, result)
	}
}

// PutNamespace
// @Tags Namespace
// @Summary 修改一个命名空间
// @Router /api/v1/namespace/:id [PUT]
func PutNamespace(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	body := models.Namespace{}
	g.ShouldBindJSON(&body)

	result, err := models.UpdateNamespace(id, &body)
	if err != nil {
		g.JSON(http.StatusBadRequest, "Namespace名称不能重复")
	} else {
		g.JSON(http.StatusOK, result)
	}
}

// DeleteNamespace
// @Tags Namespace
// @Summary 删除一个命名空间
// @Router /api/v1/namespace/:id [DELETE]
func DeleteNamespace(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	num, err := models.DeleteNamespace(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	} else {
		g.JSON(http.StatusOK, fmt.Sprintf("受影响的行数：%v", num))
	}
}
