package main

import "fmt"

func main() {
	fmt.Println("Total: ", add(1, 2))

	fmt.Println("Rectangle Area: ", rectangle(20, 30))

	var area, parameter int
	area, parameter = rectangle1(40, 50)
	fmt.Println("Area: ", area)
	fmt.Println("Parameter: ", parameter)

	fmt.Println(funcArea(20, 40))

	// Closure function
	// A closure is a function value that references variables from outside its body.
	// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	l := 30
	b := 40
	func() {
		area := l * b
		fmt.Println("Area: ", area)
	}()
}

// Simple function
func add(a, b int) int {
	return a + b
}

// Simple function with return value
func rectangle(l, b int) (area int) {
	parameter := 2 * (l + b)
	fmt.Println("\nParameter:", parameter)

	area = l * b
	fmt.Println("Area:", area)

	return
}

// Multiple return values
func rectangle1(l, b int) (area, parameter int) {
	parameter = 2 * (l + b)
	fmt.Println("\nParameter:", parameter)

	area = l * b
	fmt.Println("Area:", area)

	return
}

// Anonymous function
// Anonymous functions are functions without a name.
// They are declared using the func keyword, followed by the function parameters and the function body.
var (
	funcArea = func(l, b int) int {
		return l * b
	}
)
