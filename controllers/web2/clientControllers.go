package web2

import (
	"GoMessage/models"
	"encoding/json"
	"fmt"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

//钉钉相关结构体
type ClientDingtalkUrl struct {
	Url string `json:"url"`
}
type ClientDingtalk struct {
	*models.ClientDingtalk
	RobotUrlList []ClientDingtalkUrl `json:"robot_url"`
}

//飞书相关结构体
type ClientFeishuUrl struct {
	Url string `json:"url"`
}
type ClientFeishu struct {
	*models.ClientFeishu
	RobotUrlList []ClientFeishuUrl `json:"robot_url"`
}

// 客户端接口
type Clients struct {
	beego.Controller
}

// get方法
// @router / [get]
func (c *Clients) GetAll() {
	type Resp struct {
		ResponseData []*models.Client `json:"response_data"`
	}

	list, err := models.ListClient()
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(200)
		c.Data["json"] = Resp{ResponseData: list}
		c.ServeJSON()
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Data["json"] = Resp{ResponseData: list}
	c.ServeJSON()
}

// get方法
// @router /ns/:namespace [get]
func (c *Clients) Get() {
	type Resp struct {
		ResponseData []*models.Client `json:"response_data"`
	}

	ns := c.Ctx.Input.Param(":namespace")
	nsObject := models.GetNamespaceParamName(ns)

	list, err := models.ListNsClient(nsObject)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(200)
		c.Data["json"] = Resp{ResponseData: list}
		c.ServeJSON()
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Data["json"] = Resp{ResponseData: list}
	c.ServeJSON()
}

// post方法
// @router / [post]
func (c *Clients) Post() {

	//返回值
	type Response struct {
		Result bool `json:"result"`
	}

	//===================
	//请求的数据
	//===================
	type Param struct {
		RequestData struct {
			*models.Client                 //这里只会拿到四个数据
			ClientInfo     json.RawMessage `json:"client_info"` //这里拿到的事一个不定格式的json
		} `json:"request_data"`
	}
	param := Param{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.Data["json"] = Response{Result: false}
		c.ServeJSON()
		return
	}

	//判断客户端类型
	if param.RequestData.Type == "dingtalk" {
		dingtalk := ClientDingtalk{}
		err = json.Unmarshal(param.RequestData.ClientInfo, &dingtalk) //这里只能拿到url的绑定
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.Data["json"] = Response{Result: false}
			c.ServeJSON()
			return
		}
		var robotUrlList []string
		for _, v := range dingtalk.RobotUrlList {
			robotUrlList = append(robotUrlList, v.Url)
		}
		dingtalk.ClientDingtalk.RobotUrlList = robotUrlList
		param.RequestData.Client.ClientDingtalk = dingtalk.ClientDingtalk

	} else if param.RequestData.Type == "wechat" {
		wechat := models.ClientWechat{}
		err = json.Unmarshal(param.RequestData.ClientInfo, &wechat)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.Data["json"] = Response{Result: false}
			c.ServeJSON()
			return
		}
		param.RequestData.Client.ClientWechat = &wechat

	} else if param.RequestData.Type == "feishu" {
		feishu := ClientFeishu{}
		err = json.Unmarshal(param.RequestData.ClientInfo, &feishu)
		if err != nil {
			c.Ctx.ResponseWriter.WriteHeader(400)
			c.Data["json"] = Response{Result: false}
			c.ServeJSON()
			return
		}
		var robotUrlList []string
		for _, v := range feishu.RobotUrlList {
			robotUrlList = append(robotUrlList, v.Url)
		}
		//fmt.Println(len(robotUrlList))
		feishu.ClientFeishu.RobotUrlList = robotUrlList
		param.RequestData.Client.ClientFeishu = feishu.ClientFeishu

	} else {
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.Data["json"] = Response{Result: false}
		c.ServeJSON()
		return
	}

	//添加数据
	_, err = models.AddClient(param.RequestData.Client)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(200)
		c.Data["json"] = Response{Result: false}
		c.ServeJSON()
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Data["json"] = Response{Result: true}
	c.ServeJSON()
}

// 单个client客户端操作的控制器
type Client struct {
	beego.Controller
}

// get方法
// @router /info/:id [get]
func (c *Client) Get() {
	type RespData struct {
		*models.Client
		ClientInfo interface{} `json:"client_info"`
	}
	type Resp struct {
		ResponseData *RespData `json:"response_data"`
	}

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.Data["json"] = Resp{ResponseData: nil}
		c.ServeJSON()
		return
	}

	client, err := models.GetClient(id)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(200)
		c.Data["json"] = Resp{ResponseData: nil}
		c.ServeJSON()
		return
	}

	respData := RespData{Client: client}
	if client.Type == "dingtalk" {
		clientDingtalkUrl := []ClientDingtalkUrl{}
		for _, v := range client.ClientDingtalk.RobotUrlList {
			clientDingtalkUrl = append(clientDingtalkUrl, ClientDingtalkUrl{Url: v})
		}
		clientDingtalk := ClientDingtalk{
			ClientDingtalk: client.ClientDingtalk,
			RobotUrlList:   clientDingtalkUrl,
		}
		respData.ClientInfo = clientDingtalk

	} else if client.Type == "wechat" {
		client.ClientWechat.Secret = client.ClientWechat.Secret[:15] + "***"
		respData.ClientInfo = client.ClientWechat

	} else if client.Type == "feishu" {
		var clientFeishuUrl []ClientFeishuUrl
		for _, v := range client.ClientFeishu.RobotUrlList {
			clientFeishuUrl = append(clientFeishuUrl, ClientFeishuUrl{Url: v})
		}
		clientFeishu := ClientFeishu{
			ClientFeishu: client.ClientFeishu,
			RobotUrlList: clientFeishuUrl,
		}
		respData.ClientInfo = clientFeishu
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Data["json"] = Resp{ResponseData: &respData}
	c.ServeJSON()
}

// 删除方法
// @router /del/:id [delete]
func (c *Client) Delete() {
	type RespData struct {
		Result bool `json:"result"`
	}
	type Resp struct {
		ResponseData RespData `json:"response_data"`
	}

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.Data["json"] = Resp{ResponseData: RespData{Result: false}}
		c.ServeJSON()
		return
	}

	_, err = models.DelClient(id)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(200)
		c.Data["json"] = Resp{ResponseData: RespData{Result: false}}
		c.ServeJSON()
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Data["json"] = Resp{ResponseData: RespData{Result: true}}
	c.ServeJSON()
}

// 激活状态的控制器
type ClientActive struct {
	beego.Controller
}

// put方法
// @router /active/ [put]
func (c *ClientActive) Put() {
	type Response struct {
		Result bool `json:"result"`
	}

	type Param struct {
		Id     int  `json:"id"`
		Active bool `json:"client_active"`
	}
	param := Param{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &param)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.Data["json"] = Response{Result: false}
		c.ServeJSON()
		return
	}

	_, err = models.ActiveClient(param.Id, param.Active)
	if err != nil {
		fmt.Println(err)
		c.Ctx.ResponseWriter.WriteHeader(200)
		c.Data["json"] = Response{Result: false}
		c.ServeJSON()
		return
	}

	c.Ctx.ResponseWriter.WriteHeader(200)
	c.Data["json"] = Response{Result: true}
	c.ServeJSON()
}
