package list

import (
	. "go-practice/datastructure/base"
)

type DNode struct {
	data Object
	prev *DNode
	next *DNode
}

type DList struct {
	head *DNode
	tail *DNode
	size uint64
}

func (dList *DList) Init() {
	_dList := *(dList)
	_dList.head = nil // 没车头
	_dList.tail = nil // 没车尾
	_dList.size = 0   // 没车厢
}

func (dList *DList) Append(data Object) {
	newNode := new(DNode)
	(*newNode).data = data

	if (*dList).GetSize() == 0 { // 买个车头
		(*dList).head = newNode
		(*dList).tail = newNode
		(*newNode).prev = nil
		(*newNode).next = nil
	} else { // 排在车队尾部
		(*newNode).prev = (*dList).tail
		(*newNode).next = nil
		(*((*dList).tail)).next = newNode
		(*dList).tail = newNode
	}

	(*dList).size++
}

// 在节点后面插入数据
func (dList *DList) InsertNext(element *DNode, data Object) bool {
	if element == nil {
		return false
	}

	if dList.isTail(element) { // 恰好在尾部
		dList.Append(data)
	} else {
		newNode := new(DNode)
		(*newNode).data = data
		(*newNode).prev = element
		(*newNode).next = (*element).next

		(*element).next = newNode
		(*((*newNode).next)).prev = newNode
		(*dList).size++
	}
	return true
}

// 在节点前面插入数据，可以理解为在当前节点前一个节点的后面插入数据
func (dList *DList) InsertPrev(element *DNode, data Object) bool {
	if element == nil {
		return false
	}

	if dList.isHead(element) { // 如果是新增一个车头就特殊处理
		newNode := new(DNode)

		(*newNode).data = data
		(*newNode).prev = nil
		(*newNode).next = dList.GetHead()

		(*(dList.head)).prev = newNode
		dList.head = newNode
		(*dList).size++
		return true
	} else {
		prev := (*element).prev
		return dList.InsertNext(prev, data)
	}
}

func (dList *DList) Remove(element *DNode) Object {
	if element == nil {
		return false
	}

	prev := (*element).prev
	next := (*element).next

	if dList.isHead(element) {
		dList.head = next
	} else {
		(*prev).next = next
	}

	if dList.isTail(element) {
		dList.tail = prev
	} else {
		(*next).prev = prev
	}

	dList.size--

	return (*element).GetData()
}

func (dList *DList) Search(data Object, yourMatch ...MatchFun) *DNode {
	if dList.GetSize() == 0 {
		return nil
	}

	match := defaultMatch
	if len(yourMatch) > 0 {
		match = yourMatch[0]
	}

	node := dList.GetHead()
	for ; node != nil; node = node.GetNext() {
		if match(node.GetData(), data) == 0 {
			break
		}
	}

	return node
}

func (dList *DList) GetSize() uint64 {
	return (*dList).size
}

func (dList *DList) GetHead() *DNode {
	return (*dList).head
}

func (dList *DList) GetTail() *DNode {
	return (*dList).tail
}

func (dNode *DNode) GetData() Object {
	return (*dNode).data
}

func (dNode *DNode) GetNext() *DNode {
	return (*dNode).next
}

func (dNode *DNode) GetPrev() *DNode {
	return (*dNode).prev
}

func (dList *DList) isHead(element *DNode) bool {
	return dList.GetHead() == element
}

func (dList *DList) isTail(element *DNode) bool {
	return dList.GetTail() == element
}
