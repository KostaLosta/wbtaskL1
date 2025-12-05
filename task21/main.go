package main

import (
	"fmt"
	"strconv"
)

// OldPrinter - это старая структура
type OldPrinter struct{}

// PrintString - метод старой структуры, принимает строку
func (p *OldPrinter) PrintString(text string) {
	fmt.Println("Печать:", text)
}

// PrinterInterface - это новый интерфейс, который ожидает число
type PrinterInterface interface {
	PrintNumber(num int)
}

// PrinterAdapter - это адаптер, который решает проблему несовместимости
type PrinterAdapter struct {
	oldPrinter *OldPrinter // встроенная старая структура
}

// NewPrinterAdapter создает адаптер
func NewPrinterAdapter(oldPrinter *OldPrinter) *PrinterAdapter {
	return &PrinterAdapter{
		oldPrinter: oldPrinter,
	}
}

// PrintNumber - это метод нового интерфейса, который реализует адаптер
func (a *PrinterAdapter) PrintNumber(num int) {
	text := strconv.Itoa(num)
	a.oldPrinter.PrintString(text)
}

func main() {

	// создаем старую структуру
	oldPrinter := &OldPrinter{}

	// создаем адаптер, передавая ему старую структуру
	adapter := NewPrinterAdapter(oldPrinter)

	// клиент работает с новым интерфейсом через адаптер
	fmt.Println("Клиент передает число через новый интерфейс:")
	adapter.PrintNumber(42)

	// старая структура все еще работает напрямую (без адаптера)
	fmt.Println("Старая структура все еще работает напрямую (без адаптера):")
	oldPrinter.PrintString("Hello, World!")

}
