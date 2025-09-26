package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// worker обрабатывает данные из канала, при этом каждый воркер получает контекст для отслеживания сигнала отмены,
func worker(ctx context.Context, id int, wg *sync.WaitGroup, dataChannel <-chan string) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-dataChannel:
			// Если канал закрыт, воркер завершает работу
			if !ok {
				fmt.Printf("Воркер %d: канал закрыт, завершение работы\n", id)
				return
			}
			fmt.Printf("Воркер %d: прочел %s из канала\n", id, data)
			time.Sleep(1500 * time.Millisecond)
		case <-ctx.Done():
			// Получен сигнал отмены — завершаем работу
			fmt.Printf("Воркер %d: получен сигнал отмены\n", id)
			return
		}
	}
}

func main() {
	// Контекст используется для централизованного оповещения всех горутин о завершении.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Гарантирует освобождение ресурсов при завершении main

	// Создаем канал для получения сигналов ОС (SIGINT/SIGTERM).
	shutdown := make(chan os.Signal, 1)

	// Регистрируем обработку сигналов в shutdown.(Ctrl+C)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

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
	fmt.Println("Нажмите Ctrl+C для завершения")

	// Создаем буферизованный канал для передачи данных воркерам.
	dataChannel := make(chan string, 10)
	var wg sync.WaitGroup

	// Запускаем воркеров в отдельных горутинах, в которые передаем контекст для возможной отмены
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg, dataChannel)
	}

	// Запускаем горутину для отправки данных в канал, в горутине имеется ctx.Done() для отмены
	go func() {
		counter := 1
		for {
			select {
			case <-ctx.Done():
				// Получен сигнал отмены, завершаем работу
				fmt.Println("получен сигнал отмены")
				return
			default:
				message := fmt.Sprintf("Сообщение N%d", counter)
				select {
				case dataChannel <- message:
					// Сообщение успешно отправлено в канал
					fmt.Printf("В канал отправлено: %s\n", message)
					counter++
				case <-ctx.Done():
					// Получен сигнал отмены во время попытки отправки — завершаем работу
					fmt.Println("получен сигнал отмены при отправке")
					return
				}
				// Короткая пауза между отправками сообщений
				time.Sleep(1500 * time.Millisecond)
			}
		}
	}()

	// Ожидаем получения сигнала завершения от ОС (Ctrl+C).
	<-shutdown
	fmt.Println("Получен сигнал завершения")

	// Вызываем cancel() — все горутины, использующие ctx, получают сигнал отмены.
	cancel()

	// Короткая пауза для обработки сигнала отмены воркерами и горутиной отправки данных в канал
	time.Sleep(100 * time.Millisecond)

	// Закрываем канал данных, чтобы воркеры завершили работу после обработки всех сообщений.
	close(dataChannel)

	// Ожидаем завершения всех воркеров через WaitGroup.
	wg.Wait()

	fmt.Println("Все горутины завершили работу. Программа завершена корректно.")
}
