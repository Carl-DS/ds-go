package main

import "fmt"

func main() {
	// 和数组不同的是，切片的长度是可变的。
	// 我们可以使用内置函数make来创建一个长度不为零的切片
	// 这里我们创建了一个长度为3，存储字符串的切片，切片元素
	// 默认为零值，对于字符串就是""。
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 可以使用和数组一样的方法来设置元素值或获取元素值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get:", s[2])

	// 可以使用内置函数len获取切片的长度
	fmt.Println("len:", len(s))

	// 切片还拥有一些数组所没有的功能。
	// 例如我们可以使用内置函数append给切片追加值，然后
	// 返回一个拥有新切片元素的切片。
	// 注意append函数不会改变原切片，而是生成了一个新切片，
	// 我们需要用原来的切片来接收这个新切片
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append: ", s)
}
