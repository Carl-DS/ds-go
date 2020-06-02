package main

import (
	"fmt"
)

type T struct {
	n int
}

// map[key]struct中struct是不可寻址的，所以无法直接赋值
func main() {
	m := make(map[int]T)
	fmt.Println(m)
	//m[0].n = 1
	t := T{1}
	m[0] = t
	fmt.Println(m[0].n)
}
