package main

import (
	"fmt"
)

func main() {
	// 创建2个管道
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	go func() {
		int_chan <- 1
	}()

	go func() {
		string_chan <- "hello world"
	}()

	select {
	case value := <- int_chan:
		fmt.Println("int=", value)
	case value := <- string_chan:
		fmt.Println("string=", value)
	}

	fmt.Println("main结束")
}