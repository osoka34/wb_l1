package main

import (
	"fmt"
	"math"
)

// Определение интерфейса
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Реализация интерфейса для круга
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Реализация интерфейса для прямоугольника
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Функция, принимающая интерфейс Shape
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}

	PrintShapeInfo(circle)
	PrintShapeInfo(rectangle)
}

//Интерфейс - синтаксическая структура, определяющая отнощения между объектами,
//которые разделяют определенное поведение и не связаны никак иначе

//Абстракция - создание абстракных типов данных, которые имеют только методы, но не реализацию

//Полиморфизм - использовые одинаковых интерфейсов для разных типов данных + в гибкость кода

//Композиция - объединение нескольких интерфейсов в один

//Инкапсуляция - сокрытие данынх от пользователя, прячем детали реализации таким образом

//В Go нет наследования, но есть встраивание типов

//Как сказал Даниэль Подольский - "Единственная легитимная цель примененния интерфейсов в Go -
//это уменьшение степени связанности АКА зацепления АКА coupling вашей программы"
