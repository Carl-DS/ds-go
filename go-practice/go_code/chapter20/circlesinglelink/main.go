package main

import (
	"fmt"
)

// 定义猫的结构体结点
type CatNode struct {
	no   int // 猫猫的编号
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	// 判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head // 构成一个环形
		fmt.Println(newCatNode, "　加入到环形链表")
		return
	}

	// 定义一个临时变量，帮忙，找到环形的最后节点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	// 加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

//输出这个环形的链表
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的情况如下：")
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的环形链表...")
		return
	}
	for {
		fmt.Printf("猫的信息为［id＝%d name=%s] -> \n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {
	// 这里我们初始化一个环形链表的头节点
	head := &CatNode{}

	// 创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)
}
