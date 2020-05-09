package main

import (
	"fmt"
)

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
	fmt.Println(c.ShowA(), c.ShowB())
	fmt.Println("如果是调用没有的就会报错")
	// fmt.Println(a.ShowB())
	// fmt.Println(b.ShowA())
}