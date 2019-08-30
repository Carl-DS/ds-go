package slice

import (
	"fmt"
	"testing"
)

// 传统的数据长度固定,所以实际用途并不多,除非你明确知道自己想要多长的数组,
// 很多时候我们需要的是一个可以改变长度大小的数组,在Go 里面这类型被称为切片
func TestSlice(t *testing.T) {
	var a []int
	a = append(a, 2)
	a = append(a, 1)
	a = append(a, 4)
	a = append(a, 5)

	fmt.Printf("%v\n", a)

	a2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	s1 := a2[0:]

	s2 := a2[1:5]

	s3 := a2[4:6]

	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s3)
}
