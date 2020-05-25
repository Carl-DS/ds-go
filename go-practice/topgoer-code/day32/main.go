package main

import (
	"fmt"
)

type Foo struct {
	bar string
}

// for range 使用短变量声明(:=)的形式迭代变量时，变量i,value在每次
// 循环体中都会被重用，而不是重新声明。所以s2每次填充的都是临时变量value的地址，而
// 在最后一次循环中，value被赋值为{C}。因此，s2 输出的时候显示出了三个 &{c}。
//可行的解决办法如下：
// for i, value := range s1 {
//    s2[i] = &s1[i]
// }
func main() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		s2[i] = &value
	}
	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])
}
