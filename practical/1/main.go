package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

//Наследования нет, есть только встраивание структур внуть другой, по сути аналог наследования.

// Структура Human с произвольными полями и методами
type Human struct {
	Name  string
	Age   int
	Email string
}

func (h *Human) Speak() {
	fmt.Printf("Привет, я %s, мне %d лет.\n", h.Name, h.Age)
}

// Структура Action, встраивающая структуру Human
type Action struct {
	// Встроенная структура Human
	Human
	ActionName string
}

func main() {
	// Создаем экземпляр структуры Action
	a := Action{
		Human:      Human{Name: "Саша", Age: 100},
		ActionName: "Бег",
	}

	// Метод Speak структуры Human доступен через встраивание
	a.Speak()

	// Доступ к полям Human
	fmt.Printf("%s выполняет действие: %s\n", a.Name, a.ActionName)
}
