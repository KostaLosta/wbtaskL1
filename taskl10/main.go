package main

import "fmt"

// Функция определяет номер группы для температуры
func getGroupNumber(temp float64) int {
	// Делим температуру на 10 и округляем в меньшую сторону
	// Пример: -25.4 / 10 = -2.54 → int(-2.54) = -2 → -2 * 10 = -20
	return int(temp/10) * 10
}

func main() {
	// Исходные температуры
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Карта для хранения групп температур, де ключ номер группы, а значение слайс температур
	groups := make(map[int][]float64)

	// Проходим по всем температурам
	for _, temp := range temperatures {
		// Вычисляем номер группы для текущей температуры
		groupNumber := getGroupNumber(temp)

		// Добавляем температуру в соответствующую группу
		groups[groupNumber] = append(groups[groupNumber], temp)
	}

	// Выводим результат
	fmt.Println("Группировка температур:")
	for groupNumber, temps := range groups {
		fmt.Printf("Группа %d: %v\n", groupNumber, temps)
	}
}
