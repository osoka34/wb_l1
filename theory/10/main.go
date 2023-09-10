package main

import "fmt"

//Что выведет данная программа и почему?

func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	//1
	update(p)
	fmt.Println(*p)
	//1
}

//Мы передаем в функцию update копию указателя p, поэтому изменения внутри функции не влияют на значение p в main.
