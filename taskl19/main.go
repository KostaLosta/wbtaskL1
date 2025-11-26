package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// функция преобразует строку в []rune, чтобы корректно переворачивать Unicode.
func changeRunes(text string) string {
	runes := []rune(text) // преобразуем строку в []rune
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left] // меняем местами символы
	}
	return string(runes) // преобразуем []rune обратно в строку
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")

	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	line = strings.TrimRight(line, "\r\n") // убираем перевод строки перед разворотом.
	fmt.Println(changeRunes(line))
}
