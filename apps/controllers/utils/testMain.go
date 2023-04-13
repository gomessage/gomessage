/*
模拟组装2台电脑
--- 抽象层 ---
抽象1层
有数据渲染接口  方法（变量映射、内容渲染、直接返回原始数据）
有消息组装接口  方法（钉钉消息体、飞书消息体、微信消息体）
又消息推送接口  方法（普通的post推送、特殊的推送方法）

抽象2层
Memory接口有两个实现：（IntelMemory 和 KingstonMemory）之间的关系
数据渲染接口应该也有两个实现：（钉钉数据渲染、飞书数据渲染）

--- 实现层 ---
方式1
有 数据渲染 、产品有(变量映射、内容渲染、直接返回原始数据)
有 消息体组装， 产品有(钉钉消息体、飞书消息体、微信消息体)
有 消息体推送， 产品有(推送向钉钉、推送向飞书，推送向微信)

方式2
有 钉钉渲染推送 、产品有(变量映射，内容渲染，消息体组装、推送消息、记录器)
有 钉钉直接过境 、产品有(直接拿到原始数据、推送消息、记录器)

有 飞书渲染推送， 产品有(变量映射、内容渲染、消息体组装、推送消息、记录器)
又 飞书直接过境， 产品有(直接拿到原始数据、推送消息、记录器)

--- 逻辑层 ---
1. 向钉钉推送消息（渲染推送）
2. 向飞书推送消息（过境推送）
3. 向微信推送消息（渲染推送、过境推送）
*/

package main

//
//import (
//	"gomessage/apps/controllers/utils/base"
//	"gomessage/apps/controllers/utils/realized"
//)
//
//func main() {
//	var client2 *base.Action
//
//	client2 = base.NewAction(&realized.GetRendersResult{}, &realized.DingtalkMessageAssembled{}, &realized.GeneralPush{}, &realized.GeneralRecord{})
//	client2.Working()
//
//	client2 = base.NewAction(&realized.GeneralRequestData{}, &realized.NotRenders{}, &realized.DingtalkMessageAssembled{}, &realized.GeneralPush{}, &realized.GeneralRecord{})
//	client2.Working()
//}
