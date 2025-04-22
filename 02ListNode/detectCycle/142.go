package main

import (
	utils "carlRecommend/Common"
	"fmt"
)

func detectCycle(head *utils.ListNode) *utils.ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			index1 := fast
			index2 := head
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index1
		}
	}
	return nil
}

func main() {
	node1 := &utils.ListNode{Val: 1}
	node2 := &utils.ListNode{Val: 2}
	node3 := &utils.ListNode{Val: 3}
	node4 := &utils.ListNode{Val: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	node := detectCycle(node1)
	if node != nil {
		fmt.Printf("存在环，入口在%d\n", node.Val)
	} else {
		fmt.Println("不存在环")
	}
}
