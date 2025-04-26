package main

import (
	utils "carlRecommend/Common"
	"fmt"
)

func sortedSquares(nums []int) []int {
	i, j, k := 0, len(nums)-1, len(nums)-1
	res := make([]int, len(nums))
	for i <= j {
		if utils.Abs(nums[i]) <= utils.Abs(nums[j]) {
			res[k] = nums[j] * nums[j]
			j--
		} else {
			res[k] = nums[i] * nums[i]
			i++
		}
		k--
	}
	return res
}

func main() {
	nums := []int{-4, -2, -1, 0, 1, 2, 3}
	fmt.Println(nums)
	fmt.Println(sortedSquares(nums))
}
