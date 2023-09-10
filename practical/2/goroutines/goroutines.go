package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел,
//взятых из массива (2_version,4,6,8,10), \и выведет их квадраты в stdout.

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for i, _ := range arr {
		go func(i int) {
			defer wg.Done()
			fmt.Println(arr[i] * arr[i])
		}(i)
	}
	wg.Wait()
}
