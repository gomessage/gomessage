/*
模拟组装2台电脑
--- 抽象层 ---
有显卡Card  方法display
有内存Memory 方法storage
有处理器CPU   方法calculate

--- 实现层 ---
有 Intel因特尔公司 、产品有(显卡、内存、CPU)
有 Kingston 公司， 产品有(内存3)
有 NVIDIA 公司， 产品有(显卡)

--- 逻辑层 ---
1. 组装一台Intel系列的电脑，并运行
2. 组装一台 Intel CPU  Kingston内存 NVIDIA显卡的电脑，并运行
*/
package main

import "fmt"

/*

抽象层

*/

// Card 显卡
type Card interface {
	Display()
}

// Memory 内存
type Memory interface {
	Storage()
}

// CPU 处理器
type CPU interface {
	Calculate()
}

// Computer 电脑
type Computer struct {
	cpu  CPU
	mem  Memory
	card Card
}

func NewComputer(cpu CPU, memory Memory, card Card) *Computer {
	return &Computer{
		cpu:  cpu,
		mem:  memory,
		card: card,
	}
}

func (c *Computer) DoWork() {
	c.cpu.Calculate()
	c.mem.Storage()
	c.card.Display()
}

/*

实现层

*/

type IntelCPU struct {
	CPU
}

func (i *IntelCPU) Calculate() {
	fmt.Println("Intel CPU 开始计算了....")
}

type IntelMemory struct {
	Memory
}

func (i *IntelMemory) Storage() {
	fmt.Println("Intel Memory 开始存储了...")
}

type IntelCard struct {
	Card
}

func (i *IntelCard) Display() {
	fmt.Println("Intel Card 开始显示了...")
}

/*===================================*/

type KingstonMemory struct {
	Memory
}

func (k *KingstonMemory) Storage() {
	fmt.Println("Kingston memory storage...")
}

type NvidiaCard struct {
	Card
}

func (n *NvidiaCard) Display() {
	fmt.Println("Nvidia card display...")
}

/*

业务逻辑层

*/

func main() {

	var com3 *Computer

	com3 = NewComputer(&IntelCPU{}, &IntelMemory{}, &IntelCard{})
	com3.DoWork()

	//杂牌子
	com3 = NewComputer(&IntelCPU{}, &KingstonMemory{}, &NvidiaCard{})
	com3.DoWork()
}
