package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func test1() {
	fmt.Println("hello world打印一次")
}

func main() {
	crontabRule := "@every 1s"

	c := cron.New()
	c.Start()
	defer c.Stop()
	fmt.Println("cron启动")

	time.Sleep(time.Second * 5)
	fmt.Println("暂停5秒钟")

	c.AddFunc(crontabRule, test1)
	fmt.Println("添加定时任务")

	time.Sleep(time.Second * 20)
	fmt.Println("暂停20秒")

}
