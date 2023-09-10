package main

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
//с использованием конкурентных вычислений.

//Буферизованный и не буферизованный канал

func main() {

	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	//wg := sync.WaitGroup{}
	sum := 0

	//wg.Add(len(arr))

	for _, v := range arr {
		go func(v int) {
			//defer wg.Done()
			ch <- v * v
		}(v)
	}

	for i := 0; i < len(arr); i++ {
		sum += <-ch
	}

	//for v := range ch {
	//	sum += v
	//}

	println(sum)
}
