package main

import "fmt"

func getSum(n int) int {
	sum := 0
	for n > 0 {
		tmp := n % 10
		sum += tmp * tmp
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	uset := make(map[int]bool)
	sum := 0
	for {
		sum = getSum(n)
		if sum == 1 {
			return true
		} else if uset[sum] {
			return false
		}
		uset[sum] = true
		n = sum
	}
}

func main() {
	n := 2
	fmt.Println(isHappy(n))
}
