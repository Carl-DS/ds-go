package main

import "sync"

type SingletonSim struct {
}

var InstanceSim *SingletonSim

var mu sync.Mutex

/**
在Go语言中有个基础对象sync.Mutex，可以实现协程之间的同步逻辑
*/
func GetInstanceSim() *SingletonSim {
	mu.Lock()
	defer mu.Unlock()

	if InstanceSim == nil {
		InstanceSim = &SingletonSim{}
	}
	return InstanceSim
}
