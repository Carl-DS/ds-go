package main

import "fmt"

type MyInt1 int
type MyInt2 = int

func (i MyInt1) m1() {
	fmt.Println("MyInt1.m1")
}

//func (i MyInt2) m2(){
//	fmt.Println("MyInt2.m2")
//}

type User struct {
}
type MyUser1 User
type MyUser2 = User

func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}

func (i MyUser2) m2() {
	fmt.Println("MyUser2.m2")
}

func main() {
	var i1 MyInt1
	//var i2 MyInt2
	i1.m1()
	//i2.m2()

	var u1 MyUser1
	var u2 MyUser2
	u1.m1()
	u2.m2()

	// MyUser2完全等价于User，所以为MyUser2定义方法，等于就为User定义了方法, 接口也如此
	var i User
	i.m2()
}
