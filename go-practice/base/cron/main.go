package main

import (
	"github.com/robfig/cron"
	"log"
)

func main() {
	i := 0
	cron := cron.New()

	spec := "0/5 * * * * *"
	cron.AddFunc(spec, func() {
		i++
		log.Println("execute per five second", i)
	})
	cron.Start()
	// 注意select的用法：
	//golang 的 select 的功能和 select, poll, epoll 相似， 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。
	select {}
}
