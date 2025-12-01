package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// reverseWords разбивает строку на слова, разворачивает их порядок и склеивает обратно в строку.
func reverseWords(line string) string {
	words := strings.Fields(line) // strings.Fields разбивает строку на слова (учитывает Unicode).

	// Разворачиваем срез слов: меняем местами элементы с двух концов.
	for left, right := 0, len(words)-1; left < right; left, right = left+1, right-1 {
		words[left], words[right] = words[right], words[left]
	}

	// Склеиваем слова обратно в строку через пробел.
	return strings.Join(words, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin) // Создаем reader для чтения из stdin.
	fmt.Print("Введите строку: ")

	line, err := reader.ReadString('\n') // Читаем строку до символа новой строки.
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}
	fmt.Println(reverseWords(line))
}
