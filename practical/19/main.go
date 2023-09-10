package main

import "fmt"

//Разработать программу,
//которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.

func revert(str string) string {

	runes := []rune(str)
	out := make([]rune, len(runes))
	for i, r := range runes {
		out[len(runes)-1-i] = r
	}
	return string(out)
}

func main() {
	str := "главрыба-главрыба"
	fmt.Println(revert(str))
}
