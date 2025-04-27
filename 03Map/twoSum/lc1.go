package main

import "fmt"

func twoSum(nums []int, target int) []int {
	umap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		if idx, ok := umap[tmp]; ok {
			return []int{idx, i}
		}
		umap[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
