package main

import (
	utils "carlRecommend/Utils"
	"fmt"
	"github.com/TonyNee/carlRecommend/Utils"
)

func removeElements(head *utils.ListNode, val int) *utils.ListNode {
	dumm := new(utils.ListNode)
	dumm.Next = head
	cur := dumm
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dumm.Next
}

func main() {
	// 示例输入链表：1->2->6->3->4->5->6
	head := utils.BuildList([]int{1, 2, 6, 3, 4, 5, 6})
	fmt.Println("原始链表：")
	utils.PrintList(head)

	// 删除所有值为 6 的节点
	val := 6
	newHead := removeElements(head, val)

	fmt.Printf("删除值为 %d 的节点后：\n", val)
	utils.PrintList(newHead)
}
