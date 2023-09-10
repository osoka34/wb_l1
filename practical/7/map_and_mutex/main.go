package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.RWMutex{}
	m := make(map[int]int)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i * i
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	for k, v := range m {
		fmt.Printf("key: %v, value: %v \n", k, v)
	}
}
