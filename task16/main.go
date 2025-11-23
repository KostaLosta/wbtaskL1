package main

import "fmt"

// quickSort реализует алгоритм быстрой сортировки
func quickSort(arr []int) []int {
	// если массив пустой или содержит один элемент, он уже отсортирован. Выход
	if len(arr) <= 1 {
		return arr
	}

	// выбираем опорный элемент в середине массива
	support := arr[len(arr)/2]

	// Создаем три среза для элементов:
	var left []int   // элементы меньше опорного
	var middle []int // элементы равные опорному
	var right []int  // элементы больше опорного

	// Распределяем элементы по трем срезам
	for _, num := range arr {
		if num < support {
			left = append(left, num)
		} else if num == support {
			middle = append(middle, num)
		} else {
			right = append(right, num)
		}
	}
	// Рекурсивно сортируем левую и правую части, а затем объединяем: left + middle + right
	return append(append(quickSort(left), middle...), quickSort(right)...)
}

func main() {

	arr1 := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Исходный массив: %v\n", arr1)
	fmt.Printf("Отсортированный: %v\n\n", quickSort(arr1))

	arr2 := []int{5, 2, 8, 1, 9, 3}
	fmt.Printf("Исходный массив: %v\n", arr2)
	fmt.Printf("Отсортированный: %v\n\n", quickSort(arr2))

}
