package main

import (
	"fmt"
	"strings"
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // fl
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}

func longestCommonPrefix(strs []string) string {
	sb := strings.Builder{}
	//base := strs[0]
	//k := 0 // index for each symbol in a separate string

	// взять первую строку за базовую и в каждой последующей сравнивать символы, что не сошлось - на каждом шаге отсекать

	return sb.String()
	//for {
	//	for i := 0; i < len(strs); i++ {
	//		if i == len(strs)-1 {
	//			return sb.String()
	//		}
	//		if strs[i][k] == strs[i+1][k] {
	//			sb.WriteByte(strs[i][k])
	//			k++
	//		} else {
	//			return sb.String()
	//		}
	//	}
	//	return sb.String()
	//}
}
