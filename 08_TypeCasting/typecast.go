package main

import (
	"fmt"
	"strconv"
)

func main() {
	totalSum := 846
	number := 19

	fmt.Println("totalSum:", totalSum, "- number:", number)

	var avg float32

	// Type casting is pretty simple.
	// Example
	// v := typeName(value)
	avg = float32(totalSum) / float32(number)
	fmt.Printf("Average = %f\n", avg)

	// string to int
	s := "42"

	// strconv.Atoi() function converts a string to an integer.
	// It returns the converted integer value and an error. If the conversion fails, the error will be non-nil.
	// _ is a blank identifier. It is used to discard the error value.
	v, _ := strconv.Atoi(s)
	fmt.Println(v)

	// int to string
	i := 35

	// strconv.Itoa() function converts an integer to a string.
	str := strconv.Itoa(i)
	fmt.Println(str)
}
