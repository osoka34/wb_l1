package main

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
	// [100 2 3 4 5]
}

//append создает новый слайс, поэтому в функции someAction мы работаем с копией слайса a.
//в переданном слайсе меняем первый элемент на 100, а в копии добавляем 6 в конец.

//func someAction(v *[]int8, b int8) {
//	(*v)[0] = 100
//	*v = append(*v, b)
//}
//
//func main() {
//	var a = []int8{1, 2, 3, 4, 5}
//	someAction(&a, 6)
//	fmt.Println(a)
//}

//вот так было бы правильнее сделать, если нам всеже надо было изменить слайс в функции
