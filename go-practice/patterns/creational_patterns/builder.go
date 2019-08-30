package main

import "go-practice/patterns/creational_patterns/car"

/**
Builder patterns separates the construction of a complex object from its representation so that the same construction process can create different representations.
建造者模式：将一个复杂对象的构建与它表示分离，使得同样的构建过程可以创建不同的表示
*/

// Usage
func main() {
	assembly := car.Director{&car.ConcreteBuilder{}}
	familyCar := assembly.Builder.Color(car.BlueColor).TopSpeed(20 * car.Speed(1.2)).Build()
	familyCar.Drive()
	familyCar.Stop()

	bmwCar := assembly.Builder.Wheels(car.SteelWheels).TopSpeed(50 * car.Speed(20)).Build()
	bmwCar.Drive()
	bmwCar.Stop()
}
