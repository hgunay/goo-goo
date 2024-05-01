package main

import "fmt"

func main() {
	var a string = "Hello!"
	var b string
	c := "World!"

	fmt.Printf("Type: %T, Value: %v\n", a, a)
	fmt.Printf("Type: %T, Value: %v\n", b, b)
	fmt.Printf("Type: %T, Value: %v\n", c, c)
}
