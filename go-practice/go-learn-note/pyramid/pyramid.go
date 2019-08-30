package main

import (
	"fmt"
)

// 打印金字塔函数
func print_pyramid(totalLevel int) {

	// i 表示层数
	for i := 1; i <= totalLevel; i++ {
		//在打印*前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}

		// j 表示每层打印多少*
		//
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	print_pyramid(3)

	// 从终端输入一个整数打印出对应的金字塔
	fmt.Println("请输入打印金字塔的层数")
	var n int
	fmt.Scanln(&n)
	print_pyramid(n)
}
