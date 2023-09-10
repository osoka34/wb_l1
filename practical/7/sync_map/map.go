package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

func main() {
	m := sync.Map{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i*i)
		}(i)
	}
	wg.Wait()
	m.Range(func(key, value interface{}) bool {
		fmt.Println("Key:", key, "Value:", value)
		return true // Возвращаем true, чтобы продолжить итерацию
	})
}
