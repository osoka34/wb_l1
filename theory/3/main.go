package main

import (
	"fmt"
	"sync"
)

//Чем отличаются RWMutex от Mutex?

//RWMutex - это мьютекс, который может быть захвачен либо на чтение, либо на запись.
//Несколько горутин одновременно могут читать данные, но только одна горутина может их записывать.
//Mutex - это мьютекс, который может быть захвачен только на запись.
//Несколько горутин одновременно не могут читать и записывать данные.

func main() {
	mu := sync.Mutex{}
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