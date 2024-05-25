package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	r := 'a'

	fmt.Printf("Size: %d\n", unsafe.Sizeof(r))
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))
	fmt.Printf("Char: %c\n", r)
	fmt.Printf("Unicode: %U\n", r)

	s := "0b£"

	fmt.Printf("Unicode Points: %U\n", []rune(s))
	fmt.Println([]rune(s))
}
