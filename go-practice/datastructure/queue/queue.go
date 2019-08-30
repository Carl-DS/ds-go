package queue

import (
	. "go-practice/datastructure/base"
	. "go-practice/datastructure/list"
)

type Queue struct {
	list *List
}

// 1、Init
// 初始化队列，其实是初始化里面单链表。
func (queue *Queue) Init() {
	lst := new(List)
	(*queue).list = lst

	lst.Init()
}

// 2、Enqueue
// 小狗狗排队。
func (queue *Queue) Enqueue(data Object) bool {
	return (*queue).list.Append(data)
}

// 3、Dequeue
// 小狗狗出列。
func (queue *Queue) Dequeue() Object {
	return (*queue).list.RemoveAt(0)
}

// 4、Peek
// 时不时偷看队头狗狗
func (queue *Queue) Peek() Object {
	return (*queue).list.First()
}

// 5、GetSize
// 场地有限，队伍不能太长，得随时掌握队伍长度。
func (queue *Queue) GetSize() uint64 {
	return (*queue).list.GetSize()
}
