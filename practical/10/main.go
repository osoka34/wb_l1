package main

import (
	"fmt"
	"math"
	"sort"
)

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -10.0, -11.0, -12.0, -13.0, -14.0, -15.0}

	m := make(map[int][]float64)

	step := 10

	// Сортируем температуры
	sort.Float64s(temperatures)

	// Группируем температуры
	for _, temp := range temperatures {
		if temp < 0 {
			groupKey := (int(math.Floor(temp/float64(step))) + 1) * step
			m[groupKey] = append(m[groupKey], temp)
			continue
		}
		groupKey := int(math.Floor(temp/float64(step))) * step
		m[groupKey] = append(m[groupKey], temp)
	}

	for key, values := range m {
		fmt.Printf("%d: %v\n", key, values)
	}
}
