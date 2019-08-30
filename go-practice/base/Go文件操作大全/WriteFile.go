package main

import (
	"io/ioutil"
	"log"
)

// 快写文件
// ioutil包有一个非常有用的方法WriteFile()可以处理创建／打开文件、写字节slice和关闭文件一系列的操作。如果你需要简洁快速地写字节slice到文件中，你可以使用它。
func main() {
	err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
