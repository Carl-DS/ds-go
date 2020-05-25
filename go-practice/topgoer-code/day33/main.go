package main

import (
	"fmt"
)

func length(s string) int {
	println("call length.")
	return len(s)
}

func main() {
	s1 := "abcd"

	//for 循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量；
	//for i := 0, n :=length(s); i < n; i++ {     // 避免多次调用 length 函数。
	//	println(i, s[i])
	//}

	for i, n := 0, length(s1); i < n; i++ { // 避免多次调用 length 函数。
		println(i, s1[i])
	}

	//多重赋值分为两个步骤，有先后顺序：
	//
	//计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；
	//
	//赋值；
	//
	//所以本例，会先计算 s[i-1]，等号右边是两个表达式是常量，所以赋值运算等同于 i, s[0] = 2, "Z"。
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z"
	fmt.Printf("s: %v \n", s)
}
