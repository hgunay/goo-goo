package main

import "fmt"

func main() {
	sampleArray := [5]string{"a", "b", "c", "d", "e"}
	print(sampleArray)
}

func print(sampleArray [5]string) {
	fmt.Println(sampleArray)
}
