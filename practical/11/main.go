package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func main() {

	setA := map[int]bool{8: true, 7: true, 3: true, 4: true}
	setB := map[int]bool{4: true, 8: true, 5: true, 6: true}

	result := intersection(setA, setB)

	fmt.Println("Пересечения:", result)

}

//аналогично можно было бы и со слайсами,
//но тогда надо было бы сравнивать каждый элемент с каждым,
//что было бы не эффективно

//Быстрее слайсы было перевести в мапы и сравнивать их таким образом
//O(n) вместо O(n^2)

func intersection(set1 map[int]bool, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)
	for key := range set1 {
		if _, ok := set2[key]; ok {
			result[key] = true
		}
	}
	return result
}
