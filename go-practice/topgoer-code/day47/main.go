package main

import (
	"fmt"
)

// 2.下面代码最后一行输出什么？请说明原因。
// 知识点：变量隐藏　使用变量简短声明符号:=时，如果符号左边有多个变量，只需要保证至少有一个变量是新
// 声明的，并对已定义的变量进行赋值操作。但如果出现作用域之后，就会导致变量隐藏的问题
func main() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		i, x := 2, 2
		fmt.Println(i, x)
	}

	fmt.Println(x)
}
