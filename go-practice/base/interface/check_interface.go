package main

import (
	"fmt"
)

type x int

// 让编译器检查,确保类型实现了指定接口
// func init() {
// 	var _ fmt.Stringer = x(0)
// }

type FuncString func() string

func (f FuncString) String() string {
	return f()
}

// 字义函数类型,让相同签名的函数自动实现某个接口
func main() {
	// 对 FuncString 进行初始化再赋值给 t
	var t fmt.Stringer = FuncString(func() string {
		return "hello world!"
	})
	fmt.Println(t)

	fmt.Println(x(3))
}
