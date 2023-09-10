package main

import (
	"fmt"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

//Завершение горутины через канал

func main() {
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				time.Sleep(time.Millisecond * 100)
				fmt.Println("Hello")
			}
		}
	}()

	time.Sleep(1 * time.Second)
	done <- true

}
