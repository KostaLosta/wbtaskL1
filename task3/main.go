package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(id int, wg *sync.WaitGroup, dataChannel <-chan string) {
	defer wg.Done()

	for data := range dataChannel {
		fmt.Printf("Воркер %d: прочел %s из канала\n", id, data)
		time.Sleep(1500 * time.Millisecond)
	}

	fmt.Printf("Воркер %d завершил работу\n", id)
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	if len(os.Args) != 2 {
		fmt.Println("Использование: go run main.go <количество_воркеров>")
		fmt.Println("Пример: go run main.go 5")
		os.Exit(1)
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Ошибка: количество воркеров должно быть положительным числом")
		os.Exit(1)
	}

	fmt.Printf("Запуск %d воркеров...\n", numWorkers)

	dataChannel := make(chan string, 10)
	var wg sync.WaitGroup

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, dataChannel)
	}

	// Отдельный канал для остановки горутины записи
	stopChan := make(chan struct{})
	doneChan := make(chan struct{})

	// Горутина для записи данных
	go func() {
		defer close(doneChan)
		counter := 1
		for {
			select {
			case <-stopChan:
				fmt.Println("Остановка генерации сообщений")
				return
			default:
				message := fmt.Sprintf("Сообщение N%d", counter)
				select {
				case dataChannel <- message:
					fmt.Printf("В канал отправлено: %s\n", message)
					counter++
				case <-stopChan:
					fmt.Println("Остановка генерации сообщений")
					return
				}
				time.Sleep(1500 * time.Millisecond)
			}
		}
	}()

	// Ждем сигнала завершения
	<-sigChan
	fmt.Println("Получен сигнал завершения")

	// Останавливаем горутину записи
	close(stopChan)
	<-doneChan // Ждем пока горутина записи корректно завершится

	// Закрываем канал данных
	close(dataChannel)

	// Ждем завершения всех воркеров
	wg.Wait()

	fmt.Println("Все горутины завершили работу")
}
