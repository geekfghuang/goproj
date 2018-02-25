package main

import (
	"github.com/robfig/cron"
	"fmt"
)

func main() {
	period := "0 */1 * * * *" // 每分钟执行一次
	period = "*/2 * * * * *" //每两秒执行一次
	c := cron.New()
	c.AddFunc(period, func(){
		fmt.Printf("hello world\n")
	})
	c.Start()
	select {} // 阻塞主线程不退出
}