package main

import "fmt"

func main() {
	// Исходный слайс строк
	slice := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем карту для хранения уникальных строк, пустая структура struct{} в занчени для экономии памяти, занимает 0 байт
	set := make(map[string]struct{})

	// обход слайса
	for _, item := range slice {
		// Добавляем каждый элемент в мапу, в ней ключи уникальны
		set[item] = struct{}{}
	}

	// Выводим результат
	fmt.Println("Мапа содержит ключи:")
	for item := range set {
		fmt.Println(item)
	}

}
