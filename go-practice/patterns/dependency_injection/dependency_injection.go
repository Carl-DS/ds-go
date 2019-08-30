package main

import (
	"fmt"
	"reflect"
)

var inj *Injector

type Injector struct {
	mappers map[reflect.Type]reflect.Value // 根据类型map实际的值
}

func (inj *Injector) SetMap(value interface{}) {
	inj.mappers[reflect.TypeOf(value)] = reflect.ValueOf(value)
}

func (inj *Injector) Get(t reflect.Type) reflect.Value {
	return inj.mappers[t]
}

func (inj *Injector) Invoke(i interface{}) interface{} {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Func {
		panic("Should invoke a function!")
	}
	inValues := make([]reflect.Value, t.NumIn())
	for k := 0; k < t.NumIn(); k++ {
		inValues[k] = inj.Get(t.In(k))
	}
	ret := reflect.ValueOf(i).Call(inValues)
	return ret
}

func Host(name string, f func(a int, b string) string) {
	fmt.Println("Enter Host:", name)
	fmt.Println(inj.Invoke(f))
	fmt.Println("Exit Host:", name)
}

func Dependency(a int, b string) string {
	fmt.Println("Dependency: ", a, b)
	return `injection function exec finished ...`
}

/**
依赖注入:
具体含义是:当某个角色(可能是一个实例，调用者)需要另一个角色(另一个实例，被调用者)的协助时，在传统的程序设计过程中，通常由调用者来创建被调用者的实例。
但在这种场景下，创建被调用者实例的工作通常由容器(IoC)来完成，然后注入调用者，因此也称为依赖注入。
Golang利用函数f可以当做参数来传递，同时配合reflect包拿到参数的类型，然后根据调用者传来的参数和类型匹配上之后，最后通过reflect.Call()执行具体的函数。
*/
func main() {
	// 创建注入器
	inj = &Injector{make(map[reflect.Type]reflect.Value)}
	inj.SetMap(3030)
	inj.SetMap("zdd")

	d := Dependency
	Host("zddhub", d)

	inj.SetMap(8080)
	inj.SetMap("www.zddhub.com")
	Host("website", d)
}
