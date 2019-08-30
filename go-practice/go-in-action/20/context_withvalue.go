package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

/*
通过传递参数的方式，把name的值传递给监控函数。在这个例子里，我们实现一样的效果，但是通过的是Context的Value的方式。

我们可以使用context.WithValue方法附加一对K-V的键值对，这里Key必须是等价性的，也就是具有可比性；Value值要是线程安全的。

这样我们就生成了一个新的Context，这个新的Context带有这个键值对，在使用的时候，可以通过Value方法读取ctx.Value(key)。

记住，使用WithValue传值，一般是必须的值，不要什么值都传递。

*/
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// 附加值
	valueCtx := context.WithValue(ctx, key, "【监控1】")
	go watch2(valueCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
