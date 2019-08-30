package object_pool

import "go-practice/patterns/creational_patterns/car"

/**
The object pool creational_patterns design patterns is used to prepare and keep multiple instances according to the demand expectation.
对象池创建设计模式用于根据需求期望准备和保存多个实例。
*/
type Pool chan *car.Car

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(car.Car)
	}

	return &p
}
