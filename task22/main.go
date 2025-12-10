package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1_500_000)
	b := big.NewInt(2_500_000)

	sum := big.NewInt(0).Add(a, b)
	diff := big.NewInt(0).Sub(a, b)
	mult := big.NewInt(0).Mul(a, b)
	div := big.NewInt(0).Div(a, b)

	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("a + b = %d\n", sum)
	fmt.Printf("a - b = %d\n", diff)
	fmt.Printf("a * b = %d\n", mult)
	fmt.Printf("a / b = %d\n", div)
}
