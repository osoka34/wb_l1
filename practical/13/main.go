package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

// Вариант 1
func main() {
	var a, b int
	fmt.Println("Введите первое число:")
	fmt.Scan(&a)
	fmt.Println("Введите второе число:")
	fmt.Scan(&b)
	fmt.Println("Числа до изменения:", a, b)
	a = a + b // a = 5 + 3 = 8
	b = a - b // b = 8 - 3 = 5
	a = a - b // a = 8 - 5 = 3
	fmt.Println("Числа после изменения:", a, b)
}

// Вариант 2
//func main() {
//	var a, b int
//	fmt.Println("Введите первое число:")
//	fmt.Scan(&a)
//	fmt.Println("Введите второе число:")
//	fmt.Scan(&b)
//	fmt.Println("Числа до изменения:", a, b)
//	a, b = b, a
//	fmt.Println("Числа после изменения:", a, b)
//}
