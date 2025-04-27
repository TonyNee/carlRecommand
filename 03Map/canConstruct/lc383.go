package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	cnt := make([]int, 26)
	for _, c := range magazine {
		cnt[c-'a']++
	}
	for _, c := range ransomNote {
		cnt[c-'a']--
		if cnt[c-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	ransomNote, magazine := "abc", "cbd"
	fmt.Println(ransomNote, magazine)
	fmt.Println(canConstruct(ransomNote, magazine))
}
