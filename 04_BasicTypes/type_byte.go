package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var b1 byte = 'a'

	fmt.Printf("Size: %d\n", unsafe.Sizeof(b1))
	fmt.Printf("Type: %s\n", reflect.TypeOf(b1))
	fmt.Printf("Char: %c\n", b1)
}
