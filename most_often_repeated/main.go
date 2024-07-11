package main

import "fmt"

// Дан массив целых чисел и целое число k. Найти k наиболее часто встречающихся элементов.
// Пример
// Ввод:   nums = [1,1,1,2,2,3], k = 2
// Вывод: [1, 2]

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	result := make([]int, k)
	storage := make(map[int]int)

	// fill the map
	for _, num := range nums {
		storage[num]++
	}

	return result
}
