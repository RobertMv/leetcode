package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // 0
	fmt.Println(maxProfit([]int{2, 1, 2, 0, 1}))    // 1
	fmt.Println(maxProfit([]int{3, 2, 6, 5, 0, 3})) // 4
}

func maxProfit(prices []int) int {
	l, r, result := 0, 1, 0

	for r < len(prices) && l < r {
		a := prices[l]
		b := prices[r]
		d := b - a
		if d > result {
			result = d
		}
		if d < 0 {
			l = r
			r++
		} else {
			r++
		}
	}

	return result
}
