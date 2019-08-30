package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

// Human 对象实现SayHi 方法
func (h *Human) SayHi() {
	fmt.Println("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human 对象实现 Sing 方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

// Human 对象实现 Guzzle 方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main() {

}
