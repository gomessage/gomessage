package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomessage/pkg/models"
	"gomessage/pkg/utils"
	"net/http"
	"strconv"
)

// ListVariables
// @Tags Variables
// @Summary 获取所有用户变量
// @Router /api/v1/:namespace/vars [GET]
func ListVariables(g *gin.Context) {
	ns := g.Param("namespace")
	listVariables, err := models.ListVariables(ns)
	if err != nil {
		g.JSON(http.StatusInternalServerError, utils.ResponseFailure("数据拉取失败", err))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("数据拉取成功", listVariables))
}

// PostVariables
// @Tags Variables
// @Summary 新增一批用户变量
// @Router /api/v1/:namespace/vars [POST]
func PostVariables(g *gin.Context) {
	ns := g.Param("namespace")
	type requestData struct {
		KeyValueList []map[string]string `json:"key_value_list"`
	}
	//绑定请求数据
	body := requestData{}
	if err := g.ShouldBindJSON(&body); err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}
	ResponseVars := models.UpdateAddVars(ns, body.KeyValueList)
	g.JSON(http.StatusOK, utils.ResponseSuccessful("用户变量映射成功", ResponseVars))
}

// GetVariables
// @Tags Variables
// @Summary 查询一个用户变量
// @Router /api/v1/:namespace/vars/:id [GET]
func GetVariables(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
		return
	}
	result, err := models.GetVariablesById(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
	} else {
		g.JSON(http.StatusOK, result)
	}
}

// PutVariables
// @Tags Variables
// @Summary 修改一个用户变量
// @Router /api/v1/:namespace/vars/:id [PUT]
func PutVariables(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
		return
	}
	body := models.Variables{}
	if err = g.ShouldBindJSON(&body); err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}
	result, err := models.UpdateVariables(id, &body)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("修改失败", err))
	} else {
		g.JSON(http.StatusOK, result)
	}
}

// DeleteVariables
// @Tags Variables
// @Summary 删除一个用户变量
// @Router /api/v1/:namespace/vars/:id [DELETE]
func DeleteVariables(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("参数错误", err))
		return
	}
	num, err := models.DeleteVariables(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("删除失败", err))
	} else {
		g.JSON(http.StatusOK, fmt.Sprintf("受影响的行数：%v", num))
	}
}
