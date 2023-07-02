package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gomessage/pkg/models"
	"gomessage/pkg/utils"
	"io/ioutil"
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
			g.JSON(http.StatusInternalServerError, utils.ResponseFailure("服务器内部错误", err))
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
		return
	}
	namespace, err := models.AddNamespace(&body)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("命名空间已存在，不能重复创建", err))
	} else {

		//TODO: 如果命名空间创建成功，则自动添加一条"template"进去
		//============================
		body := models.Template{
			Namespace:    namespace.Name,
			TemplateName: namespace.Name,
		}
		// 读取 JSON 文件
		data, _ := ioutil.ReadFile("pkg/utils/template.json")
		json.Unmarshal(data, &body)
		UpdateAddTemp(namespace.Name, body)

		//TODO: 如果命名空间创建成功，则自动添加一串"变量映射"进去
		//============================
		type requestData struct {
			KeyValueList []map[string]string `json:"key_value_list"`
		}
		body2 := requestData{}
		data2, _ := ioutil.ReadFile("pkg/utils/vars.json")
		json.Unmarshal(data2, &body2)
		UpdateAddVars(namespace.Name, body2.KeyValueList)

		//给出返回值
		//============================
		g.JSON(http.StatusOK, utils.ResponseSuccessful("命名空间创建成功", &namespace))
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
	id, _ := strconv.Atoi(g.Param("id"))
	body := models.Namespace{}
	g.ShouldBindJSON(&body)

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
	id, _ := strconv.Atoi(g.Param("id"))
	ns, _ := models.GetNamespaceById(id)

	//删除变量映射
	models.DeleteVariablesByNs(ns.Name)

	//删除模板
	models.DeleteTemplateByNs(ns.Name)

	//删除客户端
	listClient, err := models.ListClient(ns.Name)
	if err != nil {
		return
	}
	for _, cli := range *listClient {
		models.DeleteClient(cli.ID)
	}

	//删除命名空间
	num, err := models.DeleteNamespace(ns.ID)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("删除失败", err))
	} else {
		g.JSON(http.StatusOK, utils.ResponseSuccessful("删除操作执行成功", fmt.Sprintf("受影响的行数：%v", num)))
	}
}
