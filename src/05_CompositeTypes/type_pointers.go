package main

import "fmt"

func main() {
	// Declaring a pointer
	var p *int

	// Assigning a value to the pointer
	i := 42
	p = &i

	// Dereferencing the pointer
	fmt.Println("Variable value: ", i) // 42
	fmt.Println("Pointer address: ", p)
	fmt.Println("Pointer value: ", *p) // 42

	// Changing the value of the pointer
	*p = 21
	fmt.Println("Pointer new value: ", *p)
	fmt.Println("Variable value: ", i) // 21
}
