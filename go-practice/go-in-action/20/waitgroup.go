package main

import (
	"fmt"
	"sync"
	"time"
)

/*
WaitGroup 是一种控制并发的方式，它的这种方式是控制多个goroutine同时完成
这种尤其适用于，好多个goroutine协同做一件事情的时候，因为每个goroutine做的都是这件事情的一部分，
只有全部的goroutine都完成，这件事情才算是完成，这是等待的方式。

实际业务场景:需要我们主动的通知某一个goroutine结束。比如我们开启一个后台goroutine一直做事情，
比如监控，现在不需要了，就需要通知这个监控goroutine结束，不然它会一直跑，就泄漏了。
*/
func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1号完成")
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2号完成")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("大家都干完了，放工")
}
