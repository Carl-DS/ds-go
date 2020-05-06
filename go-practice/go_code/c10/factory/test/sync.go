package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add(tag string) {
	for i := 0; i < 100; i++ {
		lock.Lock() // 加锁
		// fmt.Println("x = x + 1 ==== ", x, " tag==", tag, "for i == ", i)
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func main() {

	wg.Add(2)
	go add("first")
	go add("second")
	wg.Wait()
	fmt.Println(x)

}