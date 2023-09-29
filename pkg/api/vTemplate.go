package api

import (
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
	listTemplate, _ := models.ListTemplate(ns)
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
	g.ShouldBindJSON(&body)
	template := UpdateAddTemp(ns, body)
	g.JSON(http.StatusOK, utils.ResponseSuccessful("模板添加成功", &template))
}

func UpdateAddTemp(ns string, body models.Template) models.Template {
	//删除指定namespace中的模板
	listTemps, _ := models.ListTemplate(ns)
	for _, temp := range *listTemps {
		models.DeleteTemplate(temp.ID)
	}
	template, _ := models.AddTemplate(&body)
	return *template
}

// GetTemplate
// @Tags Template
// @Summary 查询一个消息模板
// @Router /api/v1/:namespace/template/:id [GET]
func GetTemplate(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	result, err := models.GetTemplateById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, "参数错误")
	} else {
		g.JSON(http.StatusOK, result)
	}
}

// PutTemplate
// @Tags Template
// @Summary 修改一个消息模板
// @Router /api/v1/:namespace/template/:id [PUT]
func PutTemplate(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	body := models.Template{}
	g.ShouldBindJSON(&body)
	result, err := models.UpdateTemplate(id, &body)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	} else {
		g.JSON(http.StatusOK, result)
	}
}

// DeleteTemplate
// @Tags Template
// @Summary 删除一个消息模板
// @Router /api/v1/:namespace/template/:id [DELETE]
func DeleteTemplate(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	num, err := models.DeleteTemplate(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	} else {
		g.JSON(http.StatusOK, fmt.Sprintf("受影响的行数：%v", num))
	}
}
