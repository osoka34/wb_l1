package main

import (
	"fmt"
	"time"
)

func Sleep(seconds float64) {
	duration := time.Duration(seconds) * time.Second
	time.Sleep(duration)
}

func main() {
	fmt.Println("Начинаем задержку на 3 секунды")
	Sleep(3)
	fmt.Println("Задержка на 3 секунды завершена")
}
