package main

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

//1 << bit создает число, в котором только i-й бит установлен в 1, а все остальные биты равны 0.
//input ^ (1 << bit) выполняет операцию XOR между input и этим числом.
//Результатом будет input, в котором i-й бит изменен на противоположное значение.

func main() {
	var input int64
	var bit int
	fmt.Println("Введите число:")
	fmt.Scan(&input)
	fmt.Println("Введите индекс бита:")
	fmt.Scan(&bit)
	fmt.Println("Число до изменения:", input)
	input = input ^ (1 << bit)
	fmt.Println("Число после изменения:", input)

}
