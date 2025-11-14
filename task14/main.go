package main

import "fmt"

func main() {
	// Функция для определения типа
	checkType := func(v interface{}) {
		// Используем switch с v.(type) для определения типа
		switch v.(type) {
		case int:
			fmt.Printf("Значение %v - это int\n", v)
		case string:
			fmt.Printf("Значение %v - это string\n", v)
		case bool:
			fmt.Printf("Значение %v - это bool\n", v)
		case chan int, chan string, chan bool:
			fmt.Printf("Значение %v - это chan\n", v)
		default:
			fmt.Printf("Значение %v - неизвестный тип\n", v)
		}
	}

	// Проверяем разные типы
	checkType(10)             // int
	checkType("hello")        // string
	checkType(true)           // bool
	checkType(make(chan int)) // chan
}
