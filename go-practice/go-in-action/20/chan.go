package main

import (
	"fmt"
	"time"
)

/*
	chan 通知
一个goroutine启动后，我们是无法控制他的，大部分情况是等待它自己结束，那么如果这个goroutine是一个不会自己结束的后台goroutine呢？比如监控等，会一直运行的。

这种情况化，一直傻瓜式的办法是全局变量，其他地方通过修改这个变量完成结束通知，然后后台goroutine不停的检查这个变量，如果发现被通知关闭了，就自我结束。

这种方式也可以，但是首先我们要保证这个变量在多线程下的安全，基于此，有一种更好的方式：chan + select 。

定义一个stop的chan，通知他结束后台goroutine。实现也非常简单，在后台goroutine中，
使用select判断stop是否可以接收到值，如果可以接收到，就表示可以退出停止了；如果没有接收到，
就会执行default里的监控逻辑，继续监控，只到收到stop的通知
*/
func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出,停止了...")
				return
			default:
				fmt.Println("goroutine 监控中")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了,通知监控停止")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

}
