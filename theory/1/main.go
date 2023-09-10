package main

import (
	"fmt"
	"strings"
)

// Какой самый эффективный способ конкатенации строк?
// Ответ: использовать пакет strings и функцию Join
func main() {
	strs := []string{"Hello", " ", "World", "!"}
	// таким образом не создается временная строка
	result := strings.Join(strs, "")
	fmt.Println(result)

	// таким образом создается временная строка
	//str1 := "Hello"
	//str2 := "World"
	//result := str1 + " " + str2
}
