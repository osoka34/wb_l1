package main

import (
	"fmt"
	"sync"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

func main() {
	// Создаем два канала: один для входных чисел и другой для результата
	inputChan := make(chan int)
	outputChan := make(chan int)

	var wg sync.WaitGroup

	// Горутина для чтения чисел из inputChan и умножения на 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range inputChan {
			outputChan <- v * 2
		}
		close(outputChan) // Закрываем outputChan после завершения
	}()

	// Горутина для вывода результата из outputChan
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range outputChan {
			fmt.Println(v)
		}
	}()

	// Входные числа
	numbers := []int{1, 2, 3, 4, 5}

	// Пишем числа в inputChan
	for _, num := range numbers {
		inputChan <- num
	}

	close(inputChan) // Закрываем inputChan, чтобы завершить горутину чтения,
	// иначе она будет ждать новых данных и произойдет deadlock

	// Ожидаем завершения горутин
	wg.Wait()
}
