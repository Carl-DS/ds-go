package main

import (
	"fmt"
)

func test() *int {
	a := 0x100
	return &a
}

/**
既然变参是切片,那么参数复制的仅是切片自身,并不包括底层数组,也因此可修改
原数据.如果需要,可用内置函数copy 复制底层数据
*/
func testChangeParam(a ...int) {
	fmt.Println(a)

	for i := range a {
		a[i] += 100
	}
}

func main() {
	fmt.Println("hello world")

	// func add(x, y int) int { //syntax error: unexpected add, expecting (
	// 	return x + y
	// }

	var a *int = test()
	println(a, *a)

	b := [3]int{1, 2, 3}
	testChangeParam(b[:]...)
	fmt.Println(b)
}
