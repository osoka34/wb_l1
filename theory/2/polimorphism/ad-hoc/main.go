package main

import "fmt"

//Полиморфизм - использовые одинаковых интерфейсов для разных типов данных + в гибкость кода

//Вообще это способность функции обрабатывать данные разных типов,
//при этом результаты ее работы будут разными в зависимости от типа данных.

//В данном случае это не истинный полиморфизм, а ad-hoc полиморфизм,
//тк ориентирован на работу с конкретным типом данных, а не с абстракциями, как истинный

type SomeInterface interface {
	SomeMethod(in string) (out string)
}

type SomeStruct struct {
	prefix string
}

func (s SomeStruct) SomeMethod(in string) (out string) {
	return s.prefix + ": " + in
}

type SomeOtherStruct struct {
	postfix string
}

func (s SomeOtherStruct) SomeMethod(in string) (out string) {
	return in + ": " + s.postfix
}

func main() {
	var iv1 SomeInterface = SomeStruct{prefix: "I say"}
	var iv2 SomeInterface = SomeOtherStruct{postfix: "I say"}

	fmt.Printf("method %q\n", iv1.SomeMethod("hello"))
	fmt.Printf("method %q\n", iv2.SomeMethod("hello"))

}
