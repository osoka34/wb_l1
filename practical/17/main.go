package main

//Реализовать бинарный поиск встроенными методами языка.

func binarySearch(sorted []int, search int) (index int) {

	if len(sorted) == 0 {
		return -1
	}
	left := 0
	right := len(sorted) - 1
	for left <= right {
		mid := (left + right) / 2
		if sorted[mid] == search {
			return mid
		} else if sorted[mid] < search {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	index := binarySearch(arr, 3)
	println(index)
}
