package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("Приложение с демонстрацией остановки горутин различными способами")

	fmt.Println("Горутина с контекстом, с завершением через вызов cancel()")
	ctx, cancel := context.WithCancel(context.Background())
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Остановка горутины через контекст")
				return
			default:
				fmt.Println("   - Горутина работает с контекстом...")
				time.Sleep(500 * time.Millisecond)
			}
		}

	}()
	time.Sleep(2 * time.Second)
	cancel()
	wg1.Wait()
	fmt.Println("________________________")

	fmt.Println("Горутина с закрытием канала")
	dataChannel := make(chan int, 10)

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for {
			data, ok := <-dataChannel
			if !ok {
				fmt.Println("Горутина завершена, канал закрыт")
				return
			}
			fmt.Printf("   - Горутина обрабатывает данные в канале %d, \n", data)
			time.Sleep(200 * time.Millisecond)
		}

	}()
	dataChannel <- 1
	dataChannel <- 2
	close(dataChannel)
	wg2.Wait()
	fmt.Println("________________________")

	fmt.Println("Горутина с выходом по условию")
	var wg3 sync.WaitGroup
	wg3.Add(1)
	go func() {
		defer wg3.Done()
		counter := 0
		for {
			if counter >= 3 {
				fmt.Println("   - Горутина завершена по условию (счетчик достиг 3)")
				return
			}
			fmt.Printf("   - Горутина работает, счетчик: %d\n", counter)
			counter++
			time.Sleep(300 * time.Millisecond)
		}

	}()
	wg3.Wait()
	fmt.Println("________________________")

	fmt.Println("Горутина с runtime.Goexit()")
	var wg4 sync.WaitGroup
	wg4.Add(1)
	go func() {
		defer wg4.Done()
		for i := 1; i <= 5; i++ {
			fmt.Printf("   - Работаю... шаг %d\n", i)
			time.Sleep(300 * time.Millisecond)

			if i == 3 {
				fmt.Println("   - Вызываю runtime.Goexit()")
				runtime.Goexit()
			}
		}

	}()
	wg4.Wait()
	fmt.Println("________________________")

	fmt.Println("Горутина с контекстом по таймауту (WithTimeout)")
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer timeoutCancel()

	var wg5 sync.WaitGroup
	wg5.Add(1)
	go func() {
		defer wg5.Done()
		for {
			select {
			case <-timeoutCtx.Done():
				fmt.Println("   - Горутина завершена контекстом (Timeout)")
				return
			default:
				fmt.Println("   - Горутина работает с таймаутом...")
				time.Sleep(700 * time.Millisecond)
			}
		}
	}()
	wg5.Wait()
	fmt.Println("________________________")

	fmt.Println("Приложение с демонстрацией остановки горутин различными способами ЗАВЕРШЕНО")
}
