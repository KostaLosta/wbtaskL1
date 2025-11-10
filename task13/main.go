package main

import "fmt"

func main() {

	fmt.Println("=== Переприсваивание значений ===")
	a1, b1 := 5, 10
	fmt.Printf("До обмена: a = %d, b = %d\n", a1, b1)
	a1, b1 = b1, a1
	fmt.Printf("После обмена: a = %d, b = %d\n", a1, b1)

	fmt.Println("\n=== Сложение/вычитание ===")
	a2, b2 := 5, 10
	fmt.Printf("До обмена: a = %d, b = %d\n", a2, b2)
	a2 = a2 + b2
	b2 = a2 - b2
	a2 = a2 - b2
	fmt.Printf("После обмена: a = %d, b = %d\n", a2, b2)

	fmt.Println("\n=== XOR-обмен ===")
	a3, b3 := 5, 10
	fmt.Printf("До обмена: a = %d, b = %d\n", a3, b3)
	a3 = a3 ^ b3
	b3 = a3 ^ b3
	a3 = a3 ^ b3
	fmt.Printf("После обмена: a = %d, b = %d\n", a3, b3)

}
