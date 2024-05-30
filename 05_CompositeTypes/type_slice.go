package main

import "fmt"

func main() {
	// Declaring a slice using make
	s := make([]string, 2, 3)
	fmt.Println(s)

	p := []string{"a", "b", "c", "d", "e"}
	fmt.Println(p)

	p = append(p, "f")
	fmt.Println(p)
}
