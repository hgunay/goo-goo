package main

import "fmt"

func main() {
	// float32: 32-bit floating-point numbers
	var x float32 = 3.14
	var y float32 = 3.4e+38

	fmt.Printf("Type: %T - Value: %v\n", x, x)
	fmt.Printf("Type: %T - Value: %v\n", y, y)

	// float64: 64-bit floating-point numbers
	var i float64 = 1.7e+308
	fmt.Printf("Type: %T - Value: %v\n", i, i)
}
