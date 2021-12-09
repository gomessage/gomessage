package prometheus

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

//控制器：接收 PrometheusLinux 版的消息推送
type LinuxController struct {
	beego.Controller
}

// @Title 接收一条AlertManager发出的webhook格式的报警信息
// @Description 接收一条AlertManager发出的webhook格式的报警信息
// @Param	body   body   body   true   "这个对象内容"
// @Success 200 {string} Success
// @Failure 403 权限不足
// @router /linux [post]
func (o *LinuxController) Post() {
	m := Messages{}
	json.Unmarshal(o.Ctx.Input.RequestBody, &m)
	o.Data["json"] = &m
	o.ServeJSON()
}
func (o *LinuxController) Get() {}

func (o *LinuxController) GetAll() {}

func (o *LinuxController) Put() {}

func (o *LinuxController) Delete() {}
