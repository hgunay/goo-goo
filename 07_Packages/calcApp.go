package main

import (
	"fmt"

	calc "07_Packages/internal"
)

func main() {
	number1 := 10
	number2 := 20

	fmt.Println("Number 1:", number1, " and Number 2:", number2)
	fmt.Println("Sum:", calc.Sum(number1, number2))
	fmt.Println("Sub:", calc.Sub(number1, number2))
	fmt.Println("Mul:", calc.Mul(number1, number2))
	fmt.Println("Div:", calc.Div(number1, number2))
}
