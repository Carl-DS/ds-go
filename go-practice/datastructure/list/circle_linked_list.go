package list

import (
	. "go-practice/datastructure/base"
)

type CNode struct {
	data Object
	next *CNode
}

type CList struct {
	size uint64
	head *CNode // 车头
}

func (cList *CList) Init() {
	lst := *cList
	lst.size = 0   // 没车厢
	lst.head = nil // 没车头
}

// 将数据添加到数据的尾部
func (cList *CList) Append(data Object) bool {
	node := new(CNode)
	(*node).data = data // 安排一个新车厢,装上data

	if cList.GetSize() == 0 {
		(*cList).head = node // 第一辆车,直接作为车头
	} else {
		item := cList.GetHead() // 找到车尾
		for ; (*item).next != cList.GetHead(); item = (*item).next {
		}
		(*item).next = node // 把新车厢挂到车尾
	}

	(*node).next = (*cList).head // 车尾再挂上车头
	(*cList).size++

	return true
}

// 在当前节点的后面,插入新的节点
func (cList *CList) InsertNext(element *CNode, data Object) bool {
	if element == nil {
		return false
	}

	node := new(CNode) // 安排一个新车厢，装上data
	(*node).data = data

	(*node).next = (*element).next // elmt后面车厢，挂在新车厢后面
	(*element).next = node         // element 后面挂上新车厢

	(*cList).size++

	return true
}

// Remove 删除结点
func (cList *CList) Remove(element *CNode) Object {
	if element == nil {
		return false
	}

	item := cList.GetHead() // 找到 element 的前面一节车厢
	for ; (*item).next != element; item = (*item).next {
	}

	(*item).next = (*element).next // 将前面一节车厢的绳索直接挂到后面一节车厢
	(*cList).size--

	return element.GetData() // 返回elmt车厢装的货物
}

// GetHead获取链表开头
func (cList *CList) GetHead() *CNode {
	return (*cList).head
}

// GetSize获取链表节点数量
func (cList *CList) GetSize() uint64 {
	return (*cList).size
}

// GetData获取节点装的数据 GetData是节点的方法，获取车厢里装的货物。
func (node *CNode) GetData() Object {
	return (*node).data
}

// GetNext获取下一个节点
// 和GetData一样是节点的方法，用于获取下一个车厢。
func (node *CNode) GetNext() *CNode {
	return (*node).next
}
