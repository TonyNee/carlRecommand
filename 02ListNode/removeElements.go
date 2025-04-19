package main

import (
	"fmt"
	"github.com/TonyNee/carlRecommend/common"
)

func removeElements(head *ListNode, val int) *ListNode {
	dumm := new(ListNode)
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
	head := buildList([]int{1, 2, 6, 3, 4, 5, 6})
	fmt.Println("原始链表：")
	printList(head)

	// 删除所有值为 6 的节点
	val := 6
	newHead := removeElements(head, val)

	fmt.Printf("删除值为 %d 的节点后：\n", val)
	printList(newHead)
}
