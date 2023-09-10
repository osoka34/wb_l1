package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := new(Person) // Создать указатель на структуру Person
	fmt.Println(p)   // &{ 0}
}
