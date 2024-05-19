package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		fmt.Println("每秒打印一次")
	})

	c.Start()
	time.Sleep(time.Second * 20)

}
