package main

import (
	"fmt"
)

func main() {
	p := Person{"张三"}
	fmt.Printf("原始Person的内存地址是：%p\n", &p)
	modifyPerson(p)
	fmt.Println(p)
}

type Person struct {
	Name string
}

func modifyPerson(p Person) {
	fmt.Printf("函数里接收到Person的内存地址是：%p\n", &p)
	p.Name = "李四"
}
