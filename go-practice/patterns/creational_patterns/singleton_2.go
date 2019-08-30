package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Instance interface {
	TestFunc()
}

type Singleton struct {
	Attr string
}

func (this *Singleton) TestFunc() {
	fmt.Println(this.Attr)
}

var (
	once     sync.Once
	instance *Singleton
)

func GetInstance(str string) *Singleton {
	once.Do(func() {
		instance = &Singleton{Attr: str}
	})
	return instance
}

/**
在程序运行过程中只产生一个实例
*/
func main() {
	for i := 0; i < 10; i++ {
		go func() {
			s := GetInstance("test: " + strconv.Itoa(i))
			s.TestFunc()
		}()
	}

	time.Sleep(1e5)
}
