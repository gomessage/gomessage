package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gomessage/pkg/models"
	"gomessage/pkg/utils"
	"net/http"
	"strconv"
)

// ListTemplate
// @Tags Template
// @Summary 获取所有消息模板
// @Router /api/v1/:namespace/template [GET]
func ListTemplate(g *gin.Context) {
	ns := g.Param("namespace")
	listTemplate, err := models.ListTemplate(ns)
	if err != nil {
		g.JSON(http.StatusInternalServerError, utils.ResponseFailure("查询失败", err))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("数据查询成功", listTemplate))
}

// PostTemplate
// @Tags Template
// @Summary 新增一个消息模板
// @Router /api/v1/:namespace/template [POST]
func PostTemplate(g *gin.Context) {
	ns := g.Param("namespace")
	body := models.Template{
		Namespace:    ns,
		TemplateName: ns,
	}
	if err := g.ShouldBindJSON(&body); err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}
	template, err := UpdateAddTemp(ns, body)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("模板添加失败", err))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("模板添加成功", &template))
}

func UpdateAddTemp(ns string, body models.Template) (models.Template, error) {
	listTemps, err := models.ListTemplate(ns)
	if err != nil {
		return models.Template{}, err
	}
	for _, temp := range *listTemps {
		if _, err := models.DeleteTemplate(temp.ID); err != nil {
			return models.Template{}, err
		}
	}
	template, err := models.AddTemplate(&body)
	if err != nil {
		return models.Template{}, err
	}
	return *template, nil
}

// GetTemplate
// @Tags Template
// @Summary 查询一个消息模板
// @Router /api/v1/:namespace/template/:id [GET]
func GetTemplate(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
		return
	}
	result, err := models.GetTemplateById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("查询失败", err))
		return
	}
	if result.Namespace != g.Param("namespace") {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", errors.New("模板不属于当前命名空间")))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("查询成功", result))
}

// PutTemplate
// @Tags Template
// @Summary 修改一个消息模板
// @Router /api/v1/:namespace/template/:id [PUT]
func PutTemplate(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
		return
	}
	exist, err := models.GetTemplateById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("查询失败", err))
		return
	}
	if exist.Namespace != g.Param("namespace") {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", errors.New("模板不属于当前命名空间")))
		return
	}
	body := models.Template{}
	if err = g.ShouldBindJSON(&body); err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}
	body.Namespace = g.Param("namespace")
	result, err := models.UpdateTemplate(id, &body)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("更新失败", err))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("修改成功", result))
}

// DeleteTemplate
// @Tags Template
// @Summary 删除一个消息模板
// @Router /api/v1/:namespace/template/:id [DELETE]
func DeleteTemplate(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
		return
	}
	exist, err := models.GetTemplateById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("查询失败", err))
		return
	}
	if exist.Namespace != g.Param("namespace") {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", errors.New("模板不属于当前命名空间")))
		return
	}
	num, err := models.DeleteTemplate(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("删除失败", err))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("删除成功", fmt.Sprintf("受影响的行数：%v", num)))
}
