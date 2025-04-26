package main

import "fmt"

func removeElement(nums []int, val int) int {
	slow := 0
	for _, v := range nums {
		if v != val {
			nums[slow] = v
			slow++
		}
	}
	return slow
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 5, 5}
	fmt.Println(nums)
	fmt.Println(removeElement(nums, 5))
	fmt.Println(nums)
}
