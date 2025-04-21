package main

import (
	utils "carlRecommend/Common"
	"fmt"
)

func reverseList(head *utils.ListNode) *utils.ListNode {
	var cur, pre *utils.ListNode = head, nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func main() {
	head := utils.BuildList([]int{1, 2, 3, 4, 5})
	fmt.Println("原始链表：")
	utils.PrintList(head)
	head2 := reverseList(head)
	fmt.Println("反转链表：")
	utils.PrintList(head2)
}
