package set

import (
	. "go-practice/datastructure/base"
	. "go-practice/datastructure/list"
)

type Set struct {
	list *List
}

// 1、Init
// 初始化集合，本质是初始化链表。
// 要比较集合中的元素，我们得传入一个比较函数，这里的match是我们的自定义类型MatchFun
func (set *Set) Init(match ...MatchFun) {
	lst := new(List)
	(*set).list = lst

	if len(match) == 0 {
		lst.Init()
	} else {
		lst.Init(match[0])
	}
}

// 2、Insert
// 把元素放入集合中。
func (set *Set) Insert(data Object) bool {
	if !set.IsMember(data) {
		return (*set).list.Append(data)
	}

	return false
}

// 3、IsEmpty
// 是否是空集合。
func (set *Set) IsEmpty() bool {
	return (*set).list.IsEmpty()
}

// 4、IsMember
// 是否是集合元素。
func (set *Set) IsMember(data Object) bool {
	return (*set).list.IsMember(data)
}

// 5、Remove
// 删除指定集合元素。
func (set *Set) Remove(data Object) bool {
	return (*set).list.Remove(data)
}

// 6、Union
//并集计算。
func (set *Set) Union(set1 *Set) *Set {
	if set1 == nil {
		return nil
	}

	nSet := new(Set)
	nSet.Init((*((*set).list)).MyMatch)

	if set.IsEmpty() && set1.IsEmpty() {
		return nSet
	}

	for i := uint64(0); i < set.getSize(); i++ {
		nSet.Insert(set.getAt(i))
	}

	var data Object
	for i := uint64(0); i < set1.getSize(); i++ {
		data = set1.getAt(i)
		if !nSet.IsMember(data) {
			nSet.Insert(data)
		}
	}
	return nSet
}

// 7、InterSection
// 计算交集。
func (set *Set) Intersection(set1 *Set) *Set {
	if set1 == nil {
		return nil
	}

	nSet := new(Set)
	nSet.Init((*((*set).list)).MyMatch)

	if set.IsEmpty() || set1.IsEmpty() {
		return nSet
	}

	fSet := set
	sSet := set1
	length := set.getSize()

	if set1.getSize() < length {
		fSet = set1
		sSet = set
	}

	var data Object
	for i := uint64(0); i < length; i++ {
		data = fSet.getAt(i)
		if sSet.IsMember(data) {
			nSet.Insert(data)
		}
	}

	return nSet
}

// 8、Difference
// 计算差集。 返回的集合是属于set，但不属于set1的集合。
func (set *Set) Difference(set1 *Set) *Set {
	if set1 == nil {
		return nil
	}

	nSet := new(Set)

	nSet.Init((*((*set).list)).MyMatch)
	if set.IsEmpty() {
		return nSet
	}

	var data Object
	for i := uint64(0); i < set.getSize(); i++ {
		data = set.getAt(i)

		if !set1.IsMember(data) {
			nSet.Insert(data)
		}
	}

	return nSet
}

// 9、IsSubSet
// 确认subSet是否是set的子集。
func (set *Set) IsSubSet(subSet *Set) bool {
	if set == nil {
		return false
	}

	if subSet == nil {
		return true
	}

	for i := uint64(0); i < subSet.getSize(); i++ {
		if !set.IsMember(subSet.getAt(i)) {
			return false
		}
	}
	return true
}

// 10、Equals
// 判断set和set1中集合元素是否一样。
func (set *Set) Equals(set1 *Set) bool {
	if set == nil || set1 == nil {
		return false
	}

	if set.IsEmpty() && set1.IsEmpty() {
		return true
	}

	nSet := set.Intersection(set1)

	return set.getSize() == nSet.getSize()
}

func (set *Set) getAt(i uint64) Object {
	return (*set).list.GetAt(i)
}

func (set *Set) getSize() uint64 {
	return (*set).list.GetSize()
}

// 11、访问集合元素
// 因为集合是没有顺序的，所以没法用序号来访问集合元素（虽然这里是用单链表实现）。这里我们用迭代器的方式来实现元素的访问。首先我们定义一个迭代器的接口。
// 因为Iterator是接口，没法保存状态，所以我们得定义一个类型来保存每次访问的游标。这里的游标是序号。
type SetIterator struct {
	index uint64
	set   *Set
}

// GetIterator
// 返回一个实现了Iterator接口的对象。
func (set *Set) GetIterator() *SetIterator {
	iterator := new(SetIterator)
	(*iterator).index = 0
	(*iterator).set = set

	return iterator
}

// HasNext
// 是否有其他元素没访问到？
func (iterator *SetIterator) HasNext() bool {
	set := (*iterator).set
	index := (*iterator).index

	return index < set.getSize()
}

// Next
// 获取其他元素。
func (iterator *SetIterator) Next() Object {
	set := (*iterator).set
	index := (*iterator).index

	if index < set.getSize() {
		data := set.getAt(index)
		(*iterator).index++

		return data
	}
	return nil
}
