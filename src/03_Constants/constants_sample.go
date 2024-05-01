package main

import "fmt"

// ConstValueA Untyped constant
const ConstValueA = "Hello, GO!"

// ConstValueB Typed constant
const ConstValueB int = 1

const (
	A int = 1
	B     = 2
	C     = "Hello, GO!"
)

func main() {
	fmt.Println("Untyped Const  value A: ", ConstValueA)
	fmt.Println("Typed Const value B: ", ConstValueB)
	fmt.Println("Multiple Const: ", A, B, C)
}
