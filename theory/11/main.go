package main

import (
	"fmt"
	"sync"
)

//Что выведет данная программа и почему?

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

//4
//1
//2
//0
//3
//fatal error: all goroutines are asleep - deadlock!

//в функцию передается копия wg, поэтому в горутине wg.Done() не уменьшает внешний wg, а только копию
//поэтому в конце цикла wg.Wait() ждет завершения 5 горутин, которые не завершаются, т.к. не уменьшают внешний wg
//поэтому происходит deadlock
//передавать надо по ссылке
//go func(wg *sync.WaitGroup, i int) {
//	fmt.Println(i)
//	wg.Done()
//}(&wg, i)
