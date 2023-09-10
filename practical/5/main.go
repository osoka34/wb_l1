package main

import (
	"context"
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 7*time.Second)
	ch := make(chan int)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
				ch <- 1
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(<-ch)
		}
	}

}
