package main

import (
	"fmt"
)

/**
和map、chan都不一样的slice

slice和map、chan都不太一样的，一样的是，它也是引用类型，它也可以在函数中修改对应的内容。

当是slice类型的时候，返回是slice这个结构体里，字段Data第一个元素的地址

所以我们通过%p打印的slice变量ages的地址其实就是内部存储数组元素的地址，slice是一种结构体+元素指针的混合类型，通过元素array(Data)的指针，可以达到修改slice里存储元素的目的。

所以修改类型的内容的办法有很多种，类型本身作为指针可以，类型里有指针类型的字段也可以。

单纯的从slice这个结构体看，我们可以通过modify修改存储元素的内容，但是永远修改不了len和cap，因为他们只是一个拷贝，如果要修改，那就要传递*slice作为参数才可以。

所以slice类型也是引用类型。
*/
func main() {
	ages := []int{6, 6, 6}
	fmt.Printf("原始slice的内存地址是: %p\n", ages)
	modifyAges(ages)
	fmt.Println(ages)
}

func modifyAges(ages []int) {
	fmt.Printf("函数里接收到slice的内存地址是: %p\n", ages)
	ages[0] = 1
	ages[1] = 2
	ages[2] = 3
}
