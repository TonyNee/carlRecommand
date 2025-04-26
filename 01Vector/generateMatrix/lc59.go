package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	l, r, t, b := 0, n-1, 0, n-1
	counter := 1
	for {
		for i := l; i <= r; i++ {
			matrix[t][i] = counter
			counter++
		}
		t++
		if t > b {
			break
		}
		for i := t; i <= b; i++ {
			matrix[i][r] = counter
			counter++
		}
		r--
		if r < l {
			break
		}
		for i := r; i >= l; i-- {
			matrix[b][i] = counter
			counter++
		}
		b--
		if b < t {
			break
		}
		for i := b; i >= t; i-- {
			matrix[i][l] = counter
			counter++
		}
		l++
		if l > r {
			break
		}
	}
	return matrix
}

func main() {
	n := 3
	matrix := generateMatrix(n)
	for _, row := range matrix {
		fmt.Println(row)
	}
}
