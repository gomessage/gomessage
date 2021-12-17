package text

import (
	"GoMessage/apps/ClientWechat"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

// 转发文本格式的数据
type WechatControllers struct {
	beego.Controller
}

// @Title 推送数据
// @Description 推送数据到微信客户端
// @Param    message    body    string    true    "对象内容"
// @Success 200 推送成功
// @Failure 403 创建失败
// @router /wechat [post]
func (this *WechatControllers) Post() {
	type Content struct {
		Message string `json:"message"`
	}
	c := Content{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &c)
	//如果出现错误，则代表参数解析或者绑定失败
	if err != nil {
		fmt.Printf("request参数解析失败")
	}
	w := wechat.WeChat{}
	bbb := w.PushMessage(c.Message)
	//返回值
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = struct {
		Result wechat.PushMessageReturn `json:"wechat_response"`
	}{Result: bbb}
	this.ServeJSON()
}
