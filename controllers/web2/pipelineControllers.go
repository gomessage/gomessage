package web2

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

//控制器：多管道入口
type PipelineControllers struct {
	beego.Controller
}

func (this *PipelineControllers) Post() {
	label := this.Ctx.Input.Param(":label")
	fmt.Println(label)

	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = label
	err := this.ServeJSON()
	if err != nil {
		return
	}
}
