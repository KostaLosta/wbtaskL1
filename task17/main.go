package main

import (
	"fmt"
	"sort"
)

// binarySearch ищет target в отсортированном срезе и возвращает индекс или -1.
func binarySearch(arr []int, target int) int {
	// sort.SearchInts выполняет бинарный поиск и возвращает индекс элемента.
	idx := sort.SearchInts(arr, target)
	if idx < len(arr) && arr[idx] == target {
		return idx
	}
	return -1
}

func main() {
	// срез для теста
	sorted := []int{-10, -3, 0, 5, 9, 12, 18}
	fmt.Println("Массив:", sorted)

	for _, target := range []int{-3, 12, 7, -10} {
		foundIdx := binarySearch(sorted, target)
		if foundIdx >= 0 {
			fmt.Printf("Элемент %d найден под индексом %d\n", target, foundIdx)
		} else {
			fmt.Printf("Элемент %d не найден\n", target)
		}
	}
}
