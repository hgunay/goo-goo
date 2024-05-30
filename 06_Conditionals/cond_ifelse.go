package main

import "fmt"

func main() {
	// if statement
	a := 6
	if a > 5 {
		fmt.Println("a is greater than 5")
	}

	b := 4
	if b > 3 && b < 5 {
		fmt.Println("b is between 3 and 5")
	}

	// if else statement
	c := 34
	if c == 34 {
		fmt.Println("İstanbul")
	} else {
		fmt.Println("Other city")
	}

	// if else if statement
	age := 20
	if age < 18 {
		fmt.Println("You are a kid")
	} else if age >= 18 && age < 40 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are old")
	}

	// if with short statement
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	}
}
