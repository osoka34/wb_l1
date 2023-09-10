package main

import (
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
//с использованием конкурентных вычислений.

//Писать даныне можно в канал, мапу, слайс, атомик, переменную, суть +- одна и та же

func main() {

	arr := []int{2, 4, 6, 8, 10}
	sum := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(len(arr))

	for _, v := range arr {
		go func(v int) {
			defer wg.Done()
			mu.Lock()
			sum += v * v
			mu.Unlock()
		}(v)
	}
	wg.Wait()

	println(sum)
}
