package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	//a и b в этом примере 2^21 и 2^22
	a.Exp(big.NewInt(2), big.NewInt(21), nil)
	b.Exp(big.NewInt(2), big.NewInt(22), nil)

	// Операция умножения
	resultMul := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %s * %s = %s\n", a.String(), b.String(), resultMul.String())

	// Операция деления
	resultDiv := new(big.Int).Div(a, b)
	fmt.Printf("Деление: %s / %s = %s\n", a.String(), b.String(), resultDiv.String())

	// Операция сложения
	resultAdd := new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %s + %s = %s\n", a.String(), b.String(), resultAdd.String())

	// Операция вычитания
	resultSub := new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание: %s - %s = %s\n", a.String(), b.String(), resultSub.String())
}
