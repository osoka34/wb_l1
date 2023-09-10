package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет,
//что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.

func Unique(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]bool)

	for _, char := range str {
		if m[char] {
			return false
		}
		m[char] = true
	}

	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(Unique(str1))
	// true
	fmt.Println(Unique(str2))
	// false
	fmt.Println(Unique(str3))
	// false
}
