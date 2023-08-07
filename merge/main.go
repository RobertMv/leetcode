package main

import "fmt"

// Дано 2 отсортированных (по возрастанию) массива A и B длины M и N.
// Нужно слить их в один отсортированный (по возрастанию) массив, состоящий из элементов первых двух.

// Пример 1
// # Ввод
// [1, 2, 5]
// [2, 2, 3, 4, 6]
// # Вывод
// [1, 2, 2, 2, 3, 4, 5, 6]

func main() {
	fmt.Println(merge([]int{1, 2, 5}, []int{2, 2, 3, 4, 6}))
}

func merge(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	if i < len(arr1) {
		result = append(result, arr1[i:]...)
	}

	if j < len(arr2) {
		result = append(result, arr2[j:]...)
	}

	return result
}
