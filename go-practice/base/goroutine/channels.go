package main

import (
	"fmt"
	"time"
)

// Channels（多个goroutines间的数据同步与通信）
func main() {

	// 创建一个channel
	c := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		// 发送数据到 channel 中
		c <- "message from closure"
	}() // 这个()表示调用该函数
	msg := <-c // 阻塞直到接收到数据
	fmt.Println(msg)
}
