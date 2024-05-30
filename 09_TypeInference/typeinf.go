package main

import "fmt"

func main() {
	name := "Hakan Gunay" // Type declaration is not needed. This is short declaration.
	fmt.Printf("Variable 'name' is of type '%T' and value '%v'\n", name, name)

	// Multiple variable declaration
	var firstName, lastName, age, salary = "Hakan", "Gunay", 40, 100000.0
	fmt.Printf("First Name: %T, Last Name: %T, Age:  %T, Salary:  %T\n", firstName, lastName, age, salary)
}
