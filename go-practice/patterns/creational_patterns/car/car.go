package car

import "fmt"

// 产品角色
type Car struct {
	Color  Color
	Wheels Wheels
	Speed  Speed
}

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportWheels Wheels = "sports"
	SteelWheels        = "steel"
)

// 建造者角色
type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Car
}

// 具体的建造者
type ConcreteBuilder struct {
	ACar Car
}

func (c *ConcreteBuilder) Color(color Color) Builder {
	c.ACar.Color = color
	return c
}

func (c *ConcreteBuilder) Wheels(wheels Wheels) Builder {
	c.ACar.Wheels = wheels
	return c
}

func (c *ConcreteBuilder) TopSpeed(speed Speed) Builder {
	c.ACar.Speed = speed
	return c
}

func (c *ConcreteBuilder) Build() Car {
	return c.ACar
}

//导演着角色
type Director struct {
	Builder Builder
}

type ICar interface {
	Drive() error
	Stop() error
}

func (car *Car) Drive() error {
	fmt.Printf("the color is %s car, on speed %v running \n", car.Color, car.Speed)
	return nil
}

func (car *Car) Stop() error {
	fmt.Printf("the color is %s car, stop \n", car.Color)
	return nil
}
