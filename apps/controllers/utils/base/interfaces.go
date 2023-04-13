package base

import (
	"gomessage/apps/controllers/send"
	"gomessage/apps/models"
)

// Renders 内容体渲染接口
type Renders interface {
	RendersData(thisNamespaceUserConfig send.NamespaceUserConfig, requestByte []byte) []string
}

// Assembled 消息体组装接口
type Assembled interface {
	AssembledData(isMerge bool, client *models.Client, contentList []string) (string, []any)
}

// Push 消息体推送接口
type Push interface {
	PushData(url string, data any)
}

// Record 记录器
type Record interface {
	RecordData()
}

// Action 发送行为
type Action struct {
	renders   Renders   //渲染内容体接口
	assembled Assembled //组装消息体接口
	push      Push      //推送数据接口
	record    Record    //记录器接口
}

// Working 行为对象的工作方法
func (c *Action) Working(requestByte []byte, thisNamespaceUserConfig send.NamespaceUserConfig, client *models.Client) {
	contentList := c.renders.RendersData(thisNamespaceUserConfig, requestByte)
	url, data := c.assembled.AssembledData(thisNamespaceUserConfig.MsgMerge, client, contentList)
	for _, dd := range data {
		c.push.PushData(url, dd)
	}

	c.record.RecordData()
}

// NewAction 创建一个新的"行为对象"
func NewAction(renders Renders, assembled Assembled, push Push, record Record) *Action {
	return &Action{
		renders:   renders,
		assembled: assembled,
		push:      push,
		record:    record,
	}
}

//type GoAction struct {
//	action    Action
//	Id        int
//	Name      string
//	Type      string
//	IsMerge   bool
//	IsRenders bool
//	Url       string
//	Data      []any
//}
//
//// Working 行为对象的工作方法
//func (c *GoAction) Working(requestByte []byte, thisNamespaceUserConfig send.NamespaceUserConfig, client *models.Client) {
//	contentList := c.action.renders.RendersData(thisNamespaceUserConfig, requestByte)
//	url, data := c.action.assembled.AssembledData(thisNamespaceUserConfig.MsgMerge, client, contentList)
//	c.action.push.PushData(url, data)
//	c.action.record.RecordData()
//}
//
//func NewGoAction(renders Renders, assembled Assembled, push Push, record Record) *Action {
//	return &Action{
//		renders:   renders,
//		assembled: assembled,
//		push:      push,
//		record:    record,
//	}
//}
