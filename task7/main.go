package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeMap - структура для безопасной работы с map в многопоточной среде
// Использует mutex для синхронизации доступа к данным
type SafeMap struct {
	mu sync.Mutex  // Мьютекс для блокировки доступа к map
	m  map[int]int // Сама map для хранения данных
}

// NewSafeMap создает новый экземпляр SafeMap с инициализированной map
func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[int]int)}
}

// Set безопасно устанавливает значение по ключу
// Блокирует доступ других горутин на время записи
func (s *SafeMap) Set(key int, value int) {
	s.mu.Lock()      // Блокируем доступ
	s.m[key] = value // Записываем значение
	s.mu.Unlock()    // Разблокируем доступ
}

// Get безопасно получает значение по ключу
// Возвращает значение и флаг наличия ключа в map
func (s *SafeMap) Get(key int) (int, bool) {
	s.mu.Lock()       // Блокируем доступ
	v, ok := s.m[key] // Читаем значение
	s.mu.Unlock()     // Разблокируем доступ
	return v, ok      // Возвращаем результат
}

// Len безопасно возвращает количество элементов в map
func (s *SafeMap) Len() int {
	s.mu.Lock()   // Блокируем доступ
	l := len(s.m) // Получаем длину
	s.mu.Unlock() // Разблокируем доступ
	return l      // Возвращаем результат
}

func main() {
	fmt.Println("Демонстрация конкурентно-безопасной записи в map (mutex)")

	// Создаем экземпляр безопасной map
	safe := NewSafeMap()

	// Параметры для тестирования
	numWorkers := 4        // Количество горутин-воркеров
	writesPerWorker := 100 // Количество записей на каждого воркера

	// WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	wg.Add(numWorkers) // Устанавливаем количество ожидаемых горутин

	// Засекаем время начала выполнения
	start := time.Now()

	// Запускаем горутины-воркеры
	for w := 0; w < numWorkers; w++ {
		workerIdx := w // Сохраняем индекс воркера для избежания замыкания
		go func() {
			defer wg.Done()                     // Уведомляем о завершении работы воркера
			base := workerIdx * writesPerWorker // Вычисляем базовый индекс для ключей

			// Каждый воркер записывает данные в свой диапазон ключей
			for i := 0; i < writesPerWorker; i++ {
				// Каждый воркер пишет в собственный диапазон ключей, чтобы видно было итоговый размер.
				safe.Set(base+i, i) // Записываем значение в безопасную map
			}
		}()
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Вычисляем время выполнения
	elapsed := time.Since(start)

	// Вычисляем ожидаемое количество элементов
	expected := numWorkers * writesPerWorker

	// Выводим результаты тестирования
	fmt.Printf("Готово. len(map)=%d, ожидается=%d, за %s\n", safe.Len(), expected, elapsed)
}
