package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	//"slices"
)

// через создание нового слайса,
// путем копирования в новый слайс всех элементов,
// кроме удаляемого
func removeElementByIndex(slice []int, index int) []int {
	// Проверяем, что индекс находится в пределах длины слайса
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Удаление элемента из слайса
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice = removeElementByIndex(slice, 2)
	fmt.Println(slice)
	//slice = [1, 2, 4, 5]

	//здесь новый слайс не создается, а изменяется существующий
	slice1 := []int{1, 2, 3, 4, 5}
	slice1 = slices.Delete(slice1, 2, 3)
	fmt.Println(slice1)
}
