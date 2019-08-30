package main

import (
	"fmt"
	"time"
)

// 通道工厂模式
func main() {
	stream := pump2()
	go suck(stream)
	time.Sleep(1e9)
}

func pump2() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	//for {
	//	fmt.Println(<-ch)
	//}
	go func() {
		for v := range ch {
			fmt.Printf("The value is %v\n", v)
		}
	}()

}
