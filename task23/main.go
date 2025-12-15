package main

import "fmt"

// removeElement удаляет i-ый элемент из слайса без утечки памяти
func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice // индекс вне диапазона
	}

	// копируем вторую часть слайса на место убираемого индекса
	copy(slice[i:], slice[i+1:])

	// обнуляем последний элемент
	slice[len(slice)-1] = 0

	// уменьшаем длину слайса на 1
	return slice[:len(slice)-1]
}

// removeElementStr удаляет i-ый элемент из слайса строк без утечки памяти
func removeElementStr(slice []string, i int) []string {
	if i < 0 || i >= len(slice) {
		return slice // индекс вне диапазона
	}

	// копируем вторую часть слайса на место убираемого индекса
	copy(slice[i:], slice[i+1:])

	// обнуляем последний элемент
	slice[len(slice)-1] = ""

	// уменьшаем длину слайса на 1
	return slice[:len(slice)-1]
}
func main() {

	sliceInt := []int{10, 20, 30, 40, 50}
	fmt.Printf("Исходный слайс: %v\n", sliceInt)

	sliceInt = removeElement(sliceInt, 2) // удаляем элемент с индексом 2 (30)
	fmt.Printf("После удаления элемента [2]: %v\n", sliceInt)

	sliceString := []string{"cat", "mouse", "dog", "bird", "lion"}
	fmt.Printf("Исходный слайс: %v\n", sliceString)

	sliceString = removeElementStr(sliceString, 2) // удаляем элемент с индексом 2 (dog)
	fmt.Printf("После удаления элемента [2]: %v\n", sliceString)
}
