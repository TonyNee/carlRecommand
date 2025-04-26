package utils

import (
	"fmt"
	"strconv"
	"strings"
)

/***************  01 Math  ***************/
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/***************  02 ListNode  ***************/
type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(nums []int) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummyHead.Next
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

/***************  03 TreeNode  ***************/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = BuildTree(nums[:mid])
	root.Right = BuildTree(nums[mid+1:])
	return root
}

// 层序打印二叉树
func PrintTreeLeveled(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Print(cur.Val, " ")
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	fmt.Printf("\n\n")
}

// 结构化打印二叉树
func PrintTreeStructured(root *TreeNode) {
	if root == nil {
		return
	}
	height := getHeight(root)
	width := (1 << height) - 1 // 2^height - 1
	res := make([][]string, height)
	for i := range res {
		res[i] = make([]string, width)
		for j := range res[i] {
			res[i][j] = " "
		}
	}
	fill(res, root, 0, 0, width-1)

	for _, row := range res {
		fmt.Println(strings.Join(row, ""))
	}
}

func fill(res [][]string, root *TreeNode, row, left, right int) {
	if root == nil {
		return
	}
	mid := (left + right) / 2
	res[row][mid] = strconv.Itoa(root.Val)
	fill(res, root.Left, row+1, left, mid-1)
	fill(res, root.Right, row+1, mid+1, right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := getHeight(root.Left)
	rh := getHeight(root.Right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}
