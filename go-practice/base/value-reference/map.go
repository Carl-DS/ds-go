package main

import (
	"fmt"
)

/**
什么是传引用(引用传递)
Go语言(Golang)是没有引用传递的，这里我不能使用Go举例子，但是可以通过说明描述。

以value.go里的例子为例，如果在modify函数里打印出来的内存地址是不变的，也是0xc42000c028，那么就是引用传递。

迷惑Map
了解清楚了传值和传引用，但是对于Map类型来说，可能觉得还是迷惑，一来我们可以通过方法修改它的内容，二来它没有明显的指针。

通过查看src/runtime/hashmap.go源代码发现，的确和我们猜测的一样，make函数返回的是一个hmap类型的指针*hmap。也就是说map===*hmap。
现在看func modify(p map)这样的函数，其实就等于func modify(p *hmap)

所以在这里，Go语言通过make函数，字面量的包装，为我们省去了指针的操作，让我们可以更容易的使用map。
这里的map可以理解为引用类型，但是记住引用类型不是传引用。

chan类型本质上和map类型是一样的，这里不做过多的介绍，参考下源代码:
func makechan(t *chantype, size int64) *hchan {
    //省略无关代码
}
chan也是一个引用类型，和map相差无几，make返回的是一个*hchan


*/
func main() {
	persons := make(map[string]int)
	persons["张三"] = 39

	mp := &persons

	fmt.Printf("原始map的内存地址是：%p\n", mp)
	modifyMap(persons)
	fmt.Println("map 值被修改了，新值为：", persons)
}

func modifyMap(p map[string]int) {
	fmt.Printf("函数里接收到map的内存地址是：%p\n", &p)
	p["张三"] = 20
}
