package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
}

func lengthOfLongestSubstring(s string) int {
	myMap := make(map[byte]struct{})
	l, res := 0, 0
	for r := range s {
		for contains(myMap, s[r]) {
			delete(myMap, s[l])
			l++
		}
		myMap[s[r]] = struct{}{}
		r++
		if r-l > res {
			res = r - l
		}
	}
	return res
}

func contains(m map[byte]struct{}, r byte) bool {
	_, ok := m[r]

	return ok
}
