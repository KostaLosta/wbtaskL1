package main

import (
	"fmt"
	"sync"
)

func main() {
	// Исходный массив чисел для обработки
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Создаем каналы для передачи данных между горутинами
	// inputChannel - для передачи исходных чисел
	// outputChannel - для передачи обработанных чисел
	inputChannel := make(chan int, len(numbers))
	outputChannel := make(chan int, len(numbers))

	// WaitGroup для синхронизации горутин
	var wg sync.WaitGroup

	// Этап 1: Горутина-генератор чисел
	// Читает числа из массива и отправляет их в inputChannel
	wg.Add(1)
	go func() {
		defer wg.Done()           // Уменьшаем счетчик WaitGroup при завершении
		defer close(inputChannel) // Закрываем канал после отправки всех чисел
		for _, num := range numbers {
			fmt.Printf("Из входящего массива в канал отправлено число %d\n", num)
			inputChannel <- num // Отправляем число в канал
		}
	}()

	// Этап 2: Горутина-обработчик
	// Читает числа из inputChannel, умножает на 2, отправляет в outputChannel
	wg.Add(1)
	go func() {
		defer wg.Done()                 // Уменьшаем счетчик WaitGroup при завершении
		defer close(outputChannel)      // Закрываем канал после обработки всех чисел
		for num := range inputChannel { // Читаем числа из inputChannel
			result := num * 2 // Умножаем на 2
			fmt.Printf("Обработчик: получил %d число, обработал и отправил результат %d в канал\n", num, result)
			outputChannel <- result // Отправляем результат в outputChannel
		}
	}()

	// Этап 3: Горутина-вывод результатов
	// Читает результаты из outputChannel и выводит в stdout
	wg.Add(1)
	go func() {
		defer wg.Done()                     // Уменьшаем счетчик WaitGroup при завершении
		for result := range outputChannel { // Читаем результаты из outputChannel
			fmt.Printf("Результат: %d\n", result) // Выводим результат в stdout
		}
	}()
	wg.Wait() // Блокируемся до завершения всех горутин
	fmt.Println("Основная горутина: конвейер завершен успешно")
}
