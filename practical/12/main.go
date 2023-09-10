package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree)
//создать для нее собственное множество.

func main() {
	set := make(map[string]bool)
	examples := []string{"cat", "cat", "dog", "cat", "tree"}

	for _, example := range examples {
		//В целом проверка здесь бессмысленна, т.к. мапа не добавит повторяющиеся значения
		if _, ok := set[example]; ok {
			continue
		}
		set[example] = true
	}

	fmt.Println(set)
}