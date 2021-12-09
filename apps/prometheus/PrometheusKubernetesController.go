package prometheus

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

//控制器：接收 PrometheusKubernetes 的信息推送
type KubernetesController struct {
	beego.Controller
}

// @Title 接收一条AlertManager发出的webhook格式的报警信息
// @Description 接收一条AlertManager发出的webhook格式的报警信息
// @Param	body   body   body   true   "这个对象内容"
// @Success 200 {string} Success
// @Failure 403 权限不足
// @router /kubernetes [post]
func (o *KubernetesController) Post() {
	m := Messages{}
	json.Unmarshal(o.Ctx.Input.RequestBody, &m)
	//apps.DisposeAlertMessage(&m)
	o.Data["json"] = &m
	o.ServeJSON()
}

func (o *KubernetesController) Get() {}

func (o *KubernetesController) GetAll() {}

func (o *KubernetesController) Put() {}

func (o *KubernetesController) Delete() {}
