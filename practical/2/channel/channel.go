package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел,
//взятых из массива (2_version,4,6,8,10), \и выведет их квадраты в stdout.

func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(arr))
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, v := range arr {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(v)
	}
	wg.Wait()

	//1_version вариант чтения из канала
	for i := 0; i < len(arr); i++ {
		fmt.Println(<-ch)
	}

	//2_version вариант чтения из канала
	//close(ch)
	//
	//for v := range ch {
	//	fmt.Println(v)
	//}

}
