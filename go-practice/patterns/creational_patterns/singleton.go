package main

import (
	"os/exec"
	"sync"
	"time"
)

var myOnlyInstance time.Time

// Go保证 init 函数会在 main 函数之前被执行，所以可以保证这些值可以在使用之前已经被初始化了
func init() {
	// Prepare things that are needed here
	//myOnlyInstance = newMyType()
	myOnlyInstance = time.Now()
}

// 如果没有初始化的前置条件,用下面的方式更好
// var myOnlyInstance = newMyType()

// 解决并发问题
var oSingle sync.Once
var single time.Time

/*
这有三个好处:
1. 保证有且只有一次调用初始化
2. 并发访问会被阻塞，直到初始化完成
3. 初始化完成后调用很快
*/
func getSingle() time.Time {
	oSingle.Do(func() {
		single = time.Now()
	})

	return single
}

/**
例如， exec.Cmd不允许调用Wait方法多次，并发访问的时候怎么办呢？你可以像这样进行包装：
*/
type multiWaitableCmd struct {
	exec.Cmd
	o   sync.Once
	err error
}

// Wait decorates `(*exec.Cmd).Wait` with a `(*sync.Once).Do()` call
func (mwc *multiWaitableCmd) Wait() error {
	mwc.o.Do(func() {
		mwc.err = mwc.Cmd.Wait()
	})
	return mwc.err
}
