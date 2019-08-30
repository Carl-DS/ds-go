package main

import (
	"fmt"
)

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

/**
闭包就是一个函数和与其相关的引用环境组合的一个整体
*/
func main() {
	fmt.Println("hello world")

	f := AddUpper()

	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
