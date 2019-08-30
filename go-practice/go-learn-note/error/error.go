package main

import (
	"fmt"
)

func test() {
	// 使用defer + recover来捕获和处理异常
	defer func() {
		// recover()内置函数,可以捕获到异常
		if err := recover(); err != nil { // 说明捕获到错误
			fmt.Println("err=", err)
			// 这里就可以将错误信息发送给管理员
			fmt.Println("发送邮件给admin@163.com")
		}
	}()
	num1 := 10
	num2 := 0

	res := num1 / num2
	fmt.Println("res=", res)
}

// 错误处理的好处
// 进行错误处理后,程序不会轻易挂掉,如果加入预警代码,就可以让程序更加的健壮
func main() {
	test()
	fmt.Println("main() 下面的代码")
}
