package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(nums []int, target int) int {
	i, sum, res := 0, 0, math.MaxInt32
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			sublen := j - i + 1
			res = min(res, sublen)
			sum -= nums[i]
			i++
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(nums)
	fmt.Println(minSubArrayLen(nums, target))
}
