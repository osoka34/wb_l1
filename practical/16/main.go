package main

import (
	"fmt"
)

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]
	var left, right []int

	for _, v := range append(arr[:pivotIndex], arr[pivotIndex+1:]...) {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quicksort(left)
	right = quicksort(right)

	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	sorted := quicksort(arr)
	fmt.Println(sorted)
}
