package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	var input int64
	ctx := context.Background()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите текст:")

	for scanner.Scan() {
		input, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		break
	}

	ch := make(chan int)

	for i := 0; i < int(input); i++ {
		go func(i int) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					time.Sleep(time.Second)
					fmt.Println(<-ch)
				}
			}
		}(i)
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- 1
		}
	}

}
