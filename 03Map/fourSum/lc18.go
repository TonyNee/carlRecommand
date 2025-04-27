package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for a := 0; a < len(nums); a++ {
		if nums[a] > 0 && nums[a] > target {
			break
		}
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < len(nums); b++ {
			if nums[a]+nums[b] > 0 && nums[a]+nums[b] > target {
				break
			}
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			c, d := b+1, len(nums)-1
			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum > target {
					d--
				} else if sum < target {
					c++
				} else {
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					for c < d && nums[c] == nums[c+1] {
						c++
					}
					c++
					for c < d && nums[d] == nums[d-1] {
						d--
					}
					d--
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{2, 2, 2, 2, 2, 2}
	res := fourSum(nums, 8)
	fmt.Println(res)
}
