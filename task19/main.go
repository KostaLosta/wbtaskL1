package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// reverseRunes преобразует строку в []rune, чтобы корректно переворачивать Unicode.
func reverseRunes(text string) string {
	runes := []rune(text) // []rune гарантирует работу с символами, а не байтами.
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left] // инвариант разворота.
	}
	return string(runes)
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
	fmt.Println(reverseRunes(line))
}
