package main

import (
	utils "carlRecommend/Common"
	"fmt"
)

func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {
	dummy := &utils.ListNode{Next: head}
	var fast, slow *utils.ListNode = dummy, dummy
	for n > 0 {
		n--
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {
	head := utils.BuildList([]int{1, 2, 3, 4, 5})
	fmt.Println("原始链表：")
	utils.PrintList(head)
	head = removeNthFromEnd(head, 2)
	fmt.Println("删除倒数第N节点：")
	utils.PrintList(head)
}
