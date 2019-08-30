package main

import "fmt"

// 结构体或者其属性的首字母大写则表示该结构体或者属性可以被导出，也就是被其它包使用。
// 结构体里面的属性成员的类型也可以是结构体，这就变相实现了类的继承
type Goods struct {
	name    string
	price   int
	pic     string
	address string
}

func main() {
	// 先定义后赋值
	var goods Goods
	goods.name = "商品1"
	goods.price = 100
	goods.pic = "http-server://xxxx.jpg"
	goods.address = "中国"

	fmt.Printf("%v\n", goods)

	// 字面量赋值
	goods2 := Goods{
		name:    "商品2",
		price:   200,
		pic:     "http-server://xxxx.png",
		address: "新加坡",
	}

	fmt.Printf("%v", goods2)
}
