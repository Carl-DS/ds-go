package main

import (
	"fmt"
	"time"
)

// Select（从多个channel中读取或写入数据）
func main() {
	fmt.Println("hello world \n")

	c1 := make(chan string)
	c2 := make(chan int8)

	go func() {
		// time.Sleep(1 * time.Second)
		c1 <- "message from closure" //发送数据到channel中
		c2 <- 3                      //发送数据到channel中
	}() //这个()表示调用该函数

	go func() {
		select {
		case v := <-c1:
			fmt.Println("channel 1 sends: \n", v)
		case v := <-c2:
			fmt.Println("channel 2 sends: \n", v)
		default: // 可选
			fmt.Println("neither channel was ready \n")
		}
	}()

	time.Sleep(5 * time.Second)
}
