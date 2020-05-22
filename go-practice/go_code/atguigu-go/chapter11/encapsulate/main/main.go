package main

import (
	"fmt"
	"go-practice/go_code/atguigu-go/chapter11/encapsulate/model"
)

func main() {

	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())

}
