package main

import "fmt"

func main() {
	// Declaring a variable with var keyword and type
	var i int = 1
	fmt.Println("With keyword and type: ", i)

	// Declaring a variable with var keyword and without type
	var j = 2
	fmt.Println("Without type: ", j)

	/*
		Declaring a variable without keyword and type
		- This is only possible inside a function for := operator
	*/
	k := 3
	fmt.Println("Without keyword and type: ", k)

	// Declaring multiple variables
	l, m := 4, 5
	fmt.Println("Multiple variables: ", l, m)

	// Variable declaration block
	var (
		a int    = 1
		b bool   = true
		c string = "Hello"
		d        = "Go"
	)
	fmt.Println("Variable block: ", a, b, c, d)

	// Value assignment after declaration
	var firstName string
	firstName = "Hakan"
	fmt.Println("First Name: ", firstName)

}
