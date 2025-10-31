package main

import "fmt"

func intersection(a, b []int) []int {
	// Создаем карту для отслеживания элементов первого множества
	set := make(map[int]bool)
	for _, num := range a {
		set[num] = true
	}

	// Собираем элементы, которые есть в обоих множествах
	var result []int
	for _, num := range b {
		if set[num] {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	result := intersection(A, B)
	fmt.Printf("Пересечение = %v\n", result) // Пересечение = [2 3]
}
