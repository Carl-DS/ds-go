package main

import "fmt"

// 串联的Channels（Pipeline）
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer(in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
