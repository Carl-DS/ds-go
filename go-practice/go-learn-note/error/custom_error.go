package main

import (
	"errors"
	"fmt"
)

// 自定义错误
// Go 程序中, 支持自定义错误,使用errors.New 和 painc 内置函数
// 1. errors.New("错误说明"), 会返回一个 error 类型的值,表示一个错误
// 2. panic 内置函数,接收一个interface{} 类型的值(也就是任何值了)作为参数,可以接收error类型的变量,输出错误信息,并退出程序
func main() {
	test()
	fmt.Println("main() 下面的代码")
}

// 函数去读取一配置文件init.conf 的信息
// 如果文件名传入不正确,我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误")
	}
}

func test() {
	err := readConf("config2.ini")
	if err != nil {
		// 如果读取文件发生错误,就输出这个错误,并终止程序
		panic(err)
	}
	fmt.Println("test() 继续执行")
}
