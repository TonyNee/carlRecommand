package main

import (
	utils "carlRecommend/Common"
	"fmt"
)

func removeElements(head *utils.ListNode, val int) *utils.ListNode {
	dummy := &utils.ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

func main() {
	head := utils.BuildList([]int{1, 2, 6, 3, 4, 5, 6})
	fmt.Println("原始链表：")
	utils.PrintList(head)

	head = removeElements(head, 6)
	fmt.Println("处理后的链表：")
	utils.PrintList(head)
}
