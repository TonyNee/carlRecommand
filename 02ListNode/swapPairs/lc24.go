package main

import (
	utils "carlRecommend/Common"
	"fmt"
)

func swapPairs(head *utils.ListNode) *utils.ListNode {
	dummyHead := &utils.ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		tmp1 := cur.Next
		tmp2 := cur.Next.Next.Next

		cur.Next = cur.Next.Next
		cur.Next.Next = tmp1
		cur.Next.Next.Next = tmp2

		cur = cur.Next.Next
	}
	return dummyHead.Next
}

func main() {
	head := utils.BuildList([]int{1, 2, 3, 4, 5})
	fmt.Println("原始链表：")
	utils.PrintList(head)
	head = swapPairs(head)
	fmt.Println("两两反转：")
	utils.PrintList(head)
}
