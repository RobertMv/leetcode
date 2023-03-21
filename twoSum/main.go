package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 1}, 9)) // [0,1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))     // [1,2]
	fmt.Println(twoSum([]int{3, 3}, 6))        // [0,1]
}

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	result := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		if v, ok := myMap[target-nums[i]]; ok {
			result[0], result[1] = v, i
			return result
		}
		myMap[nums[i]] = i
	}

	return result
}
