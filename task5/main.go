package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Использование: go run main.go <секунды>")
		os.Exit(1)
	}

	seconds, err := strconv.Atoi(os.Args[1])
	if err != nil || seconds <= 0 {
		fmt.Println("Ошибка: укажите положительное число секунд")
		os.Exit(1)
	}

	fmt.Printf("Программа запущена на %d секунд...\n", seconds)

	// Создаем канал для обмена данными.
	dataChannel := make(chan string)

	// Создаем канал завершения программы
	done := make(chan bool)

	// Горутина для отправки сообщений в канал
	go func() {
		counter := 1
		for {
			select {
			case <-done:
				close(dataChannel)
				return
			default:
				dataChannel <- fmt.Sprintf("Сообщение %d", counter)
				fmt.Printf("Сообщение %d отправлено в канал\n", counter)
				counter++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Горутина для чтения сообщений из канала
	go func() {
		for msg := range dataChannel {
			fmt.Printf("Прочитано: %s\n", msg)
		}
		fmt.Println("Канал закрыт.")
	}()

	// Таймаут с использованием time.After, по истечению которого завершается программа
	select {
	case <-time.After(time.Duration(seconds) * time.Second):
		fmt.Printf("Программа отработала %d секунд, начинается завершение программы\n", seconds)
		close(done)
	}

	// Короткая пауза для обработки
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Программа завершена")
}
