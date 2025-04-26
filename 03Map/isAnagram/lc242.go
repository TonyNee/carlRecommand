package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	umaps := make(map[rune]int)
	for _, c := range s {
		umaps[c]++
	}
	for _, c := range t {
		umaps[c]--
		if umaps[c] < 0 {
			return false
		}
	}
	return true
}

func main() {
	s, t := "absc", "bcsa"
	fmt.Println(isAnagram(s, t))
}
