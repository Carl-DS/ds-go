package main

import (
	"fmt"
	"go-practice/go_code/c10/factory/model"
)

func main() {
	// 创建要给student实例
	// var stu = model.Student{
	// 	Name: "tom",
	// 	Score: 70.9,
	// }

	var stu = model.NewStudent("tom", 82.8)

	fmt.Println(*stu)
	fmt.Println("name=>", stu.Name, " score=>", stu.Score) // 底层会处理为 (*stu).Name
}