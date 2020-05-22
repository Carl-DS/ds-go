package main

import (
	"fmt"
	"go-practice/go_code/atguigu-go/chapter11/encapexercise/model"
)

func main() {
	//创建一个account变量
	account := model.NewAccount("jzh11111", "000", 40)
	if account != nil {
		fmt.Println("创建成功=", account)
	} else {
		fmt.Println("创建失败")
	}
}
