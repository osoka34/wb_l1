package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		// создали внутри копию слайса slice, добавили в конец "a"
		slice[0] = "b"
		// изменили первый элемент на "b"
		slice[1] = "b"
		// изменили второй элемент на "b"
		fmt.Print(slice)
		// [b b a]
	}(slice)
	fmt.Print(slice)
	// исходный слайс не изменился, т.к. мы работали с копией
	// [a a]
}
