package main

import "fmt"

func change(s ...int) {
	s = append(s, 3)
}

func main() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	// Go　提供的语法糖..., 可以将slice传进可变函数，不会创建新的切片
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)

}
