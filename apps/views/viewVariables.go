package views

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomessage/apps/models"
	"net/http"
	"strconv"
)

// ListVariables
// @Tags Variables
// @Summary 获取所有用户变量
// @Router /api/v1/:namespace/vars [GET]
func ListVariables(g *gin.Context) {
	ns := g.Param("namespace")
	listVariables, _ := models.ListVariables(ns)
	g.JSON(http.StatusOK, listVariables)
}

// PostVariables
// @Tags Variables
// @Summary 新增一个用户变量
// @Router /api/v1/:namespace/vars [POST]
func PostVariables(g *gin.Context) {
	ns := g.Param("namespace")
	type requestData struct {
		KeyValueList []map[string]string `json:"key_value_list"`
	}
	r := requestData{}
	g.ShouldBindJSON(&r)

	//遍历删除当前namespace下的所有用户变量
	listVars, _ := models.ListVariables(ns)
	for _, json := range *listVars {
		models.DeleteVariables(json.ID)
	}

	var ResponseVars []models.Variables
	//批量写入新的配置
	for _, oneVar := range r.KeyValueList {
		for kk, vv := range oneVar {
			v := models.Variables{
				Namespace: ns,
				Key:       kk,
				Value:     vv,
			}
			models.AddVariables(&v)
			ResponseVars = append(ResponseVars, v)
		}
	}

	g.JSON(http.StatusOK, ResponseVars)
}

// GetVariables
// @Tags Variables
// @Summary 查询一个用户变量
// @Router /api/v1/:namespace/vars/:id [GET]
func GetVariables(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	result, err := models.GetVariablesById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, "参数错误")
	} else {
		g.JSON(http.StatusOK, result)
	}
}

// PutVariables
// @Tags Variables
// @Summary 修改一个用户变量
// @Router /api/v1/:namespace/vars/:id [PUT]
func PutVariables(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	body := models.Variables{}
	g.ShouldBindJSON(&body)

	result, err := models.UpdateVariables(id, &body)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	} else {
		g.JSON(http.StatusOK, result)
	}
}

// DeleteVariables
// @Tags Variables
// @Summary 删除一个用户变量
// @Router /api/v1/:namespace/vars/:id [DELETE]
func DeleteVariables(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	num, err := models.DeleteVariables(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	} else {
		g.JSON(http.StatusOK, fmt.Sprintf("受影响的行数：%v", num))
	}
}
