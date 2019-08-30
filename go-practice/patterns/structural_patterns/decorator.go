package main

import "fmt"

type Object func(string) string

func Decorate(fn Object) Object {
	return func(base string) string {
		ret := fn(base)

		ret = ret + " and TShirt"
		return ret
	}
}

func Dressing(cloth string) string {
	return "dressing " + cloth
}

/**
装饰模式就是在不改变对象内部结构的情况下，动态扩展它的功能。它提供了灵活的方法来扩展对象的功能。

下面是一个简单的实现逻辑,通过Decorate来进一步装饰Dressing函数：
*/
func main() {
	f := Decorate(Dressing)

	fmt.Println(f("shoes"))
}
