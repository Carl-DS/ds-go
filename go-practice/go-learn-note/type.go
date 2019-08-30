package main

import (
	"fmt"
)

type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

type ( // 组
	user struct { // 结构体
		name string
		age  uint8
	}

	event func(string) bool // 函数类型
)

func main() {
	fmt.Println("hello world")
	f := read | exec
	println(f)

	u := user{"Tom", 20}
	fmt.Println(u)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}

	f("abc")

	// 即便指定了基础类型,也只表明它们有相同底层数据结构,两者间不存在任务关系,属
	// 完全不同的两种类型
	type data int
	var d data = 10

	// var x int = d
	// printn(x)
}
