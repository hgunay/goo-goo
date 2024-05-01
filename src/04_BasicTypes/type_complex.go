package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Complex numbers are used to represent complex numbers (real and imaginary parts)
	var i complex64 = 1 + 2i
	var j complex64
	j = 2 + 3i
	k := 3 + 4i

	var x float64 = 5
	var y float64 = 6
	z := complex(x, y)

	fmt.Printf("Type: %T - Value: %v - Size: %d bytes\n", i, i, unsafe.Sizeof(i))
	fmt.Printf("Type: %T - Value: %v - Size: %d bytes\n", j, j, unsafe.Sizeof(j))
	fmt.Printf("Type: %T - Value: %v - Size: %d bytes\n", k, k, unsafe.Sizeof(k))
	fmt.Printf("Type: %T - Value: %v - Size: %d bytes\n", z, z, unsafe.Sizeof(z))
}
