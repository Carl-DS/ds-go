package main

import "fmt"

type Action interface {
	Move(int) int
}

type Animal struct {
	Name string
}

type Bird struct {
	Animal
}

func (this *Bird) Move(num int) int {
	fmt.Printf("I am a bird, I flyed %d meters.\n", num)
	return num
}

type Fish struct {
	Animal
}

func (this *Fish) Move(num int) int {
	fmt.Printf("I am a fish, I swimmed %d meters.\n", num)
	return num
}

type AnimalFactory interface {
	CreateAnimal() Action
}

type BirdFactory struct {
}

func (this *BirdFactory) CreateAnimal() Action {
	return &Bird{}
}

type FishFactory struct {
}

func (this *FishFactory) CreateAnimal() Action {
	return &Fish{}
}

/**
Factory Method 工厂方法模式：定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法使一个类的实例化延迟到其子类
*/
func main() {
	bFactory := &BirdFactory{}
	bird := bFactory.CreateAnimal()
	bird.Move(300)

	fFactory := &FishFactory{}
	fish := fFactory.CreateAnimal()
	fish.Move(300)
}
