package main

import (
	"carlRecommend/Common"
	"fmt"
)

type MyLinkedList struct {
	Size  int
	Dummy *utils.ListNode
}

func Constructor() MyLinkedList {
	dummy := &utils.ListNode{
		Val:  -1,
		Next: nil,
	}
	return MyLinkedList{
		Size:  0,
		Dummy: dummy,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.Size-1 {
		return -1
	}
	cur := this.Dummy
	for index > 0 {
		index--
		cur = cur.Next
	}
	return cur.Next.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &utils.ListNode{Val: val}
	cur := this.Dummy
	node.Next = cur.Next
	cur.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &utils.ListNode{Val: val}
	cur := this.Dummy
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	node := &utils.ListNode{Val: val}
	cur := this.Dummy
	for index > 0 {
		index--
		cur = cur.Next
	}
	node.Next = cur.Next
	cur.Next = node
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.Size-1 {
		return
	}
	cur := this.Dummy
	for index > 0 {
		index--
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	this.Size--
}

func (this *MyLinkedList) PrintList() {
	cur := this.Dummy.Next
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("nil")
}

func main() {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	obj.PrintList()
	obj.Get(1)
	obj.DeleteAtIndex(1)
	obj.PrintList()
	obj.Get(1)
}
