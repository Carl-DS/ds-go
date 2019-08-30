package main

import "fmt"

type Dog struct {
	Animal
}

func (this *Dog) Move(num int) int {
	fmt.Printf("I am a dog, I flying %d meters.\n", num)
	return num
}

type SimAnimalFactory struct {
}

func NewAnimalFactory() *SimAnimalFactory {
	return &SimAnimalFactory{}
}

func (*SimAnimalFactory) CreateAnimal(name string) Action {
	switch name {
	case "bird":
		return &Bird{}
	case "fish":
		return &Fish{}
	case "dog":
		return &Dog{}
	default:
		panic("error animal type")
		return nil
	}
}

/**
简单工厂模式（Simple Factory Pattern）：定义一个工厂类，它可以根据参数的不同返回不同类的实例，被创建的实例通常都具有共同的父类。
它属于类创建型模式。

简单工厂模式的要点在于：当你需要什么，只需要传入一个正确的参数，就可以获取你所需要的对象，而无须知道其创建细节。
*/
func main() {
	dog := NewAnimalFactory().CreateAnimal("dog")
	dog.Move(100)

	fish := NewAnimalFactory().CreateAnimal("fish")
	fish.Move(200)

	bird := NewAnimalFactory().CreateAnimal("bird")
	bird.Move(300)
}
