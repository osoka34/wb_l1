package main

import (
	"context"
	"fmt"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

// через контекст (по своей сути это канал, в котором есть еще map[string]any через которую можно передавать данные)
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Millisecond * 100)
				fmt.Println("Hello")
			}
		}
	}()

	// Когда горутина должна завершиться
	time.Sleep(1 * time.Second)
	cancel()
}
