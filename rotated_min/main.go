package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))       // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) // 0
	fmt.Println(findMin([]int{7, 0, 1, 2, 4, 5, 6})) // 0
	fmt.Println(findMin([]int{11, 13, 15, 17}))      // 11
}

func findMin(nums []int) int {
	l, r, min := 0, len(nums)-1, nums[0]

	for l <= r {
		if nums[l] < nums[r] {
			if nums[l] < min {
				min = nums[l]
				break
			}
		}
		m := (l + r) / 2
		if nums[m] < min {
			min = nums[m]
		}
		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return min
}
