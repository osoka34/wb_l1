package main

import (
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счетчик,
//которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	atmc atomic.Int64
}

//можно было бы и так, но через атомик удобнее
//type Counter struct {
//	count int
//	mu    sync.Mutex
//}

func (c *Counter) Increment() {
	c.atmc.Add(1)
}

func (c *Counter) Value() int64 {
	return c.atmc.Load()
}

func main() {
	сounter := Counter{}
	wg := sync.WaitGroup{}

	for goRoutine := 0; goRoutine < 10; goRoutine++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			сounter.Increment()
		}()
	}
	wg.Wait()

	println(сounter.Value())
}
