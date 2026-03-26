package api

import (
	"errors"
	"fmt"
	"gomessage/pkg/models"
	"gomessage/pkg/utils"
	"gomessage/pkg/utils/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
			g.JSON(http.StatusInternalServerError, utils.ResponseFailure("服务器内部错误", err))
			return
		}
		g.JSON(http.StatusOK, utils.ResponseSuccessful("查询成功", list))
	default:
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", errors.New("is_active的值只能为布尔值true、false")))
	}
}

// PostNamespace
// @Tags Namespace
// @Summary 新增一个命名空间
// @Router /api/v1/namespace [POST]
func PostNamespace(g *gin.Context) {
	body := models.Namespace{}
	if err := g.ShouldBindJSON(&body); err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}
	namespace, err := models.AddNamespace(&body)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("命名空间已存在，不能重复创建", err))
	} else {
		//如果命名空间创建成功，则自动添加一条"template"进去
		models.InitTemplate(namespace.Name)
		//如果命名空间创建成功，则自动添加一串"变量映射"进去
		models.InitVarMap(namespace.Name)
		//给出返回值
		g.JSON(http.StatusOK, utils.ResponseSuccessful("命名空间创建成功", &namespace))
	}
}

// GetNamespace
// @Tags Namespace
// @Summary 查询一个命名空间
// @Router /api/v1/namespace/:id [GET]
func GetNamespace(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
		return
	}
	result, err := models.GetNamespaceById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
	} else {
		g.JSON(http.StatusOK, utils.ResponseSuccessful("查询成功", result))
	}
}

// PutNamespace
// @Tags Namespace
// @Summary 修改一个命名空间
// @Router /api/v1/namespace/:id [PUT]
func PutNamespace(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
		return
	}
	body := models.Namespace{}
	if err = g.ShouldBindJSON(&body); err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}
	result, err := models.UpdateNamespace(id, &body)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("namespace名称不能重复", err))
	} else {
		g.JSON(http.StatusOK, utils.ResponseSuccessful("修改成功", result))
	}
}

// DeleteNamespace
// @Tags Namespace
// @Summary 删除一个命名空间
// @Router /api/v1/namespace/:id [DELETE]
func DeleteNamespace(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
		return
	}
	ns, err := models.GetNamespaceById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("查询失败", err))
		return
	}

	tx := database.DB.Default.Begin()
	if tx.Error != nil {
		g.JSON(http.StatusInternalServerError, utils.ResponseFailure("删除失败", tx.Error))
		return
	}
	if err = tx.Where("namespace = ?", ns.Name).Delete(&models.Variables{}).Error; err != nil {
		tx.Rollback()
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("删除变量映射失败", err))
		return
	}
	if err = tx.Where("namespace = ?", ns.Name).Delete(&models.Template{}).Error; err != nil {
		tx.Rollback()
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("删除模板失败", err))
		return
	}
	if err = tx.Where("namespace = ?", ns.Name).Delete(&models.Client{}).Error; err != nil {
		tx.Rollback()
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("删除客户端失败", err))
		return
	}
	result := tx.Delete(&models.Namespace{}, ns.ID)
	if result.Error != nil {
		tx.Rollback()
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("删除失败", result.Error))
		return
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("删除失败", err))
		return
	}

	num := int(result.RowsAffected)
	g.JSON(http.StatusOK, utils.ResponseSuccessful("删除操作执行成功", fmt.Sprintf("受影响的行数：%v", num)))
}
