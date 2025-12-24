package main

import (
	"fmt"
	"time"
)

// функция блокирует выполнение текущей горутины на указанное время, через создание канала и ожидание чтени из него
func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Начало выполнения")
	start := time.Now()

	// Приостанавливаем выполнение на 2 секунды
	Sleep(2 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Прошло времени: %v\n", elapsed)
	fmt.Println("Выполнение продолжено после Sleep")
}
