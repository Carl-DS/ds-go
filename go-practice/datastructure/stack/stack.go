package stack

import (
	. "go-practice/datastructure/base"
	. "go-practice/datastructure/list"
)

type Stack struct {
	list *List
}

// 1、Init
// 初始化栈，其实就是初始化里面的链表。
func (stack *Stack) Init() {
	lst := new(List)
	(*stack).list = lst
	lst.Init()
}

// 2、Push
// 数据入栈，也叫压栈，就是把车子开进去。
func (stack *Stack) Push(data Object) bool {
	lst := (*stack).list

	return lst.InsertAtHead(data) // 车子开进去
}

//3、Pop
//数据出栈，就是把车子开出来，当然是从链表头开出来了。
func (stack *Stack) Pop() Object {
	lst := (*stack).list

	return lst.RemoveAt(0) // 从链表关把车子开出来
}

// 4、Peek
// 时不时的偷看下，当前栈里的最近的车厢
func (stack *Stack) Peek() Object {
	lst := (*stack).list

	return lst.First
}

// 5、GetSize
// 要实时掌握栈里车厢数量，一旦太多，就要控制下。
func (stack *Stack) GetSize() uint64 {
	lst := (*stack).list

	return lst.GetSize()
}
