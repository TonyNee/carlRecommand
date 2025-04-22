package main

import (
	utils "carlRecommend/Common"
	"fmt"
)

func getIntersectionNode(headA *utils.ListNode, headB *utils.ListNode) *utils.ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

func main() {
	interNode := &utils.ListNode{
		Val: 4,
		Next: &utils.ListNode{
			Val: 5,
			Next: &utils.ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}
	headA := &utils.ListNode{
		Val: 1,
		Next: &utils.ListNode{
			Val: 2,
			Next: &utils.ListNode{
				Val:  3,
				Next: interNode,
			},
		},
	}
	utils.PrintList(headA)
	headB := &utils.ListNode{
		Val: 7,
		Next: &utils.ListNode{
			Val:  4,
			Next: interNode,
		},
	}
	utils.PrintList(headB)
	node := getIntersectionNode(headA, headB)
	if node == nil {
		fmt.Println("不相交")
	} else {
		fmt.Printf("相交在%d，", node.Val)
		if node.Next != nil {
			fmt.Printf("下一个节点是%d\n", node.Next.Val)
		} else {
			fmt.Println("下一个节点是nil")
		}
	}
}
