package main

import "fmt"

type Comparable[T any] interface {
	Compare(T) int
}

func FindMax[T Comparable[T]](slice []T) T {
	max := slice[0]
	for _, v := range slice {
		if v.Compare(max) > 0 {
			max = v
		}
	}
	return max
}

type IntValue int

func (iv IntValue) Compare(other IntValue) int {
	if iv > other {
		return 1
	} else if iv < other {
		return -1
	}
	return 0
}

type FloatValue float64

func (fv FloatValue) Compare(other FloatValue) int {
	if fv > other {
		return 1
	} else if fv < other {
		return -1
	}
	return 0
}

func main() {
	intSlice := []IntValue{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	maxInt := FindMax(intSlice)
	fmt.Printf("Максимальное значение в срезе intSlice: %d\n", maxInt)

	floatSlice := []FloatValue{3.14, 2.71, 1.0, 0.5, 2.0}
	maxFloat := FindMax(floatSlice)
	fmt.Printf("Максимальное значение в срезе floatSlice: %f\n", maxFloat)
}
