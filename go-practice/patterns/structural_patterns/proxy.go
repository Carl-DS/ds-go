package main

import "fmt"

// 被代理的公共函数
type ProxyFuncs interface {
	// 卖房功能
	SailHouse()
}

type MasterBeijing struct {
	Name     string // 北京业主姓名
	Location string // 业主所卖房屋的位置
}

func (this *MasterBeijing) SailHouse() {
	fmt.Printf("%s sailing house at %s\n", this.Name, this.Location)
}

type Proxier struct {
	Mofbj *MasterBeijing
}

func (this *Proxier) SailHouse() {
	if this.Mofbj == nil {
		this.Mofbj = &MasterBeijing{}
	}
	this.Mofbj.SailHouse()
}

/**
代理模式，简单来说就是提供一个对象来控制其他对象的功能。在一些情况下，一个Object不适合直接引用目标对象，但可以通过代理对象调用目标对象，起到中介代理的作用。
*/
func main() {
	m := &MasterBeijing{
		Name:     "Lao Wang",
		Location: "Xi Cheng",
	}

	proxier := &Proxier{
		Mofbj: m,
	}

	proxier.SailHouse()
}
