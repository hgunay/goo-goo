package main

import (
	"fmt"
)

func main() {
	a, b := 10, 0

	result, err := divide(a, b)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide '%d' by zero", a)
	}

	return a / b, nil
}
