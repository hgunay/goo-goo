package main

import (
	"fmt"

	"test_module/packages"
)

func main() {
	fmt.Println("Hello from module_test.go")

	packages.PrintHello()
}
