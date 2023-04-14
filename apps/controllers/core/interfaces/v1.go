package interfaces

import (
	"errors"
	"gomessage/apps/controllers/send"
	"gomessage/apps/models"
)

// Renders 内容体渲染接口
type Renders interface {
	RendersData(thisNamespaceUserConfig send.NamespaceUserConfig, requestByte []byte) []string
}

// Assembled 消息体组装接口
type Assembled interface {
	AssembledData(isRenders, isMerge bool, client *models.Client, contentList []string) (string, []any)
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

// NewAction 创建一个新的"行为对象"
func NewAction(renders Renders, assembled Assembled, push Push, record Record) *Action {
	return &Action{
		renders:   renders,
		assembled: assembled,
		push:      push,
		record:    record,
	}
}

// Working 行为对象的工作方法
func (c *Action) Working(isRenders bool, requestByte []byte, thisNamespaceUserConfig send.NamespaceUserConfig, client *models.Client) error {
	contentList := c.renders.RendersData(thisNamespaceUserConfig, requestByte)
	if len(contentList) == 0 {
		return errors.New("过境数据格式错误，用户变量无法从过境数据中找到可用的映射关系")
	}
	url, data := c.assembled.AssembledData(isRenders, thisNamespaceUserConfig.MsgMerge, client, contentList)
	for _, dd := range data {
		c.push.PushData(url, dd)
	}
	c.record.RecordData()
	return nil
}
