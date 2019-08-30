package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}
type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

// 超集接口变量可隐匿转换为子集,反过来不行
func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1

	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaInf is: %T\n", t)
	}

	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaInf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}
