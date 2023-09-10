package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func reverseWords(input string) string {
	// Разбиваем строку на слова
	words := strings.Fields(input)

	reversedWords := make([]string, len(words))
	for i, word := range words {
		reversedWords[len(words)-1-i] = word
	}
	reversedString := strings.Join(reversedWords, " ")

	return reversedString
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Println(reversed)
}
