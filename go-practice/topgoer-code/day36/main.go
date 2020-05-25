package main

import (
	"fmt"
	"runtime"
)

/**
select会随机选择一个可用通道做收发操作，所以可能触发异常，也可以不会
*/
func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
