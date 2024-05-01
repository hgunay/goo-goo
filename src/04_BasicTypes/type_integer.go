package main

import "fmt"

func main() {
	// Signed integers
	var x int = 500
	var y int = -4500

	fmt.Printf("Type: %T - Value: %v\n", x, x)
	fmt.Printf("Type: %T - Value: %v\n", y, y)

	// Unsigned integers
	var i uint = 500
	var j uint = 4500

	fmt.Printf("Type: %T - Value: %v\n", i, i)
	fmt.Printf("Type: %T - Value: %v\n", j, j)
}
