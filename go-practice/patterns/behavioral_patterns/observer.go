package main

import (
	"fmt"
	"time"
)

type (
	Event2 struct {
		Data string
	}

	Observer2 interface {
		// 更新事件
		Update(*Event2)
	}

	// 被观察的对象接口
	Subject2 interface {
		// 注册观察者
		Register(Observer2)
		// 注销观察者
		Deregister(Observer2)
		// 通知观察者事件
		Notify(Observer2)
	}

	ConcreteObserver struct {
		Id int
	}

	ConcreteSubject struct {
		Observers map[Observer2]struct{}
	}
)

func (c *ConcreteObserver) Update(e *Event2) {
	fmt.Printf("observer [%d] recieved msg: %s.\n", c.Id, e.Data)
}

func (c *ConcreteSubject) Register(ob Observer2) {
	c.Observers[ob] = struct{}{}
}

func (c *ConcreteSubject) Deregister(ob Observer2) {
	delete(c.Observers, ob)
}

func (c *ConcreteSubject) Notify(e *Event2) {
	for ob, _ := range c.Observers {
		ob.Update(e)
	}
}

/**
观察者模式简单一句话说就是当特定事件出现时，一个对象实例把事件发布到对应的观察者实例上执行相应的更新操作。
一个观察目标可以对应多个观察者，而且这些观察者之间没有相互联系，可以根据需要增加和删除观察者，
使得系统更易于扩展，这就是观察者模式的模式动机。
*/
func main() {
	cs := &ConcreteSubject{
		Observers: make(map[Observer2]struct{}),
	}

	observer1 := &ConcreteObserver{1}
	observer2 := &ConcreteObserver{2}

	cs.Register(observer1)
	cs.Register(observer2)

	for i := 0; i < 5; i++ {
		e := &Event2{fmt.Sprintf("msg [%d]", i)}
		cs.Notify(e)

		time.Sleep(time.Duration(1) * time.Second)
	}
}
