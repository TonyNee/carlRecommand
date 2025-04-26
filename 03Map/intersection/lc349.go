package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	uset := make(map[int]bool)
	for _, v := range nums1 {
		uset[v] = true
	}
	var res []int
	for _, v := range nums2 {
		if uset[v] {
			uset[v] = false
			res = append(res, v)
		}
	}
	return res
}

func main() {
	nums1, nums2 := []int{1, 2, 2, 4}, []int{2, 2, 6}
	fmt.Println(nums1, nums2)
	fmt.Println(intersection(nums1, nums2))
}
