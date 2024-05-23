package api

import (
	"fmt"
	"gomessage/pkg/models"
	"gomessage/pkg/utils"
	"net/http"
	"strconv"

	crontabJob "gomessage/pkg/crontab"

	"github.com/gin-gonic/gin"
)

// PostCrontab
// @Tags Crontab
// @Summary 新增一个Crontab任务
// @Router /api/v1/crontab [POST]
func PostCrontab(g *gin.Context) {
	var crontab models.Crontab
	if err := g.ShouldBindJSON(&crontab); err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}
	//if !utils.IsValidCrontabRule(crontab.CrontabRule) {
	//	g.JSON(http.StatusBadRequest, utils.ResponseFailure("Crontab规则格式错误", nil))
	//	return
	//}
	err := models.AddCrontab(&crontab)
	if err != nil {
		fmt.Println(err)
		g.JSON(http.StatusInternalServerError, utils.ResponseFailure("创建Crontab失败", err))
		return
	}

	if crontab.CrontabActivate {
		err := crontabJob.AddJob(crontab)
		if err != nil {
			fmt.Println(err)
			g.JSON(http.StatusInternalServerError, utils.ResponseFailure("创建Crontab失败", err))
			return
		}
	}

	g.JSON(http.StatusOK, utils.ResponseSuccessful("添加成功", crontab))
}

// ListCrontabs
// @Tags Crontab
// @Summary 获取所有Crontab任务
// @Router /api/v1/crontab [GET]
func ListCrontabs(g *gin.Context) {
	crontabs, err := models.ListCrontabs()
	if err != nil {
		g.JSON(http.StatusInternalServerError, utils.ResponseFailure("无法获取Crontab列表", err))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("获取成功", crontabs))
}

// GetCrontab
// @Tags Crontab
// @Summary 查询一个Crontab任务
// @Router /api/v1/crontab/:id [GET]
func GetCrontab(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("无效的ID", err))
		return
	}
	crontab, err := models.GetCrontab(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, utils.ResponseFailure("无法获取Crontab", err))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("获取成功", crontab))
}

// PutCrontab
// @Tags Crontab
// @Summary 修改一个Crontab任务
// @Router /api/v1/crontab/:id [PUT]
func PutCrontab(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("无效的ID", err))
		return
	}
	var crontab models.Crontab
	if err := g.ShouldBindJSON(&crontab); err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求的内容错误", err))
		return
	}
	err = models.UpdateCrontab(id, &crontab)
	if err != nil {
		g.JSON(http.StatusInternalServerError, utils.ResponseFailure("更新Crontab失败", err))
		return
	}

	crontabJob.RemoveJobByCrontabID(id)
	if crontab.CrontabActivate {
		err := crontabJob.AddJob(crontab)
		if err != nil {
			fmt.Println(err)
			g.JSON(http.StatusInternalServerError, utils.ResponseFailure("更新Crontab失败", err))
			return
		}
	}

	g.JSON(http.StatusOK, utils.ResponseSuccessful("更新成功", crontab))
}

// DeleteCrontab
// @Tags Crontab
// @Summary 删除一个Crontab任务
// @Router /api/v1/crontab/:id [DELETE]
func DeleteCrontab(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("无效的ID", err))
		return
	}
	err = models.DeleteCrontab(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, utils.ResponseFailure("删除Crontab失败", err))
		return
	}

	crontabJob.RemoveJobByCrontabID(id)

	g.JSON(http.StatusOK, utils.ResponseSuccessful("删除成功", nil))
}
