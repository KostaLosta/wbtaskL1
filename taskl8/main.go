package main

import (
	"flag"
	"fmt"
)

func main() {

	var n int64
	var i uint
	var to int

	flag.Int64Var(&n, "n", 0, "int64 value")        // читаем -n исходное число (int64)
	flag.UintVar(&i, "i", 0, "bit index (0..63)")   // читаем -i индекс бита (0-базный)
	flag.IntVar(&to, "to", 0, "bit value (0 or 1)") // читаем -to значение бита: 0 или 1
	flag.Parse()                                    // парсим флаги

	if i >= 64 {
		fmt.Println("bit index must be in [0..63]")
		return
	}
	if to != 0 && to != 1 {
		fmt.Println("to must be 0 or 1")
		return
	}
	// Формируем маску с единицей на i-й позиции.
	mask := int64(1) << i // маска с 1 в i-й позиции

	// Применяем нужную операцию и печатаем результат.
	if to == 1 { // если нужно установить бит в 1
		// Установить бит: OR поднимает соответствующий бит в 1.
		fmt.Println(n | mask) // устанавливаем бит через OR
	} else { // иначе — сбросить в 0
		// Сбросить бит: AND NOT обнуляет соответствующий бит.
		fmt.Println(n &^ mask) // сбрасываем бит через AND NOT
	}
}
