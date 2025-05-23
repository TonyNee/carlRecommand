package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	umap := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			umap[a+b]++
		}
	}
	res := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			res += umap[0-c-d]
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}
