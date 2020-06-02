package main

import (
	"fmt"
)

type T struct {
	n int
}

func (t *T) Set(n int) {
	t.n = n
}

func getT() T {
	return T{}
}

// 下面的代码有几处问题？请详细说明。
// 1. 直接返回的T{}不可寻址
// 2. 不可寻址的结构体不能调用带结构体指针接收者的方法
func main() {
	//getT().Set(2)
	// 修复
	t := getT()
	t.Set(2)
	fmt.Println(t.n)
}
