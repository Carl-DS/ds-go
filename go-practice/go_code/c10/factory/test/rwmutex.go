package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
	rwlock sync.RWMutex
)

func write(tag int) {
	// lock.Lock() // 加互斥锁
	rwlock.Lock() // 加写锁
	// fmt.Println("x = x + 1 ==== ", x, " tag==", tag)
	x += 1
	time.Sleep(time.Millisecond * 10) // 假设读操作耗时10毫秒
	// lock.Unlock() // 解互斥锁
	rwlock.Unlock() // 解写锁
	wg.Done()
}

func read() {
	// lock.Lock() // 加互斥锁
	rwlock.RLock() // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	// lock.Unlock() // 解互斥锁
	rwlock.RUnlock() // 解读锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 60; i++ {
		wg.Add(1)
		go write(i)
	}

	for i :=0; i < 2000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println("打印结果: ", x)
}