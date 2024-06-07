package main

import "fmt"

func main() {
	fmt.Println("------- Classic for loop  -------")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("Index: ", i, " Sum: ", sum)
	}

	fmt.Println("Total Sum: ", sum)

	fmt.Println("------- For loop with a single condition -------")

	// If you skip the init and post statements, you get a while loop.
	i := 1
	for i <= 3 {
		fmt.Println("Single Condition: ", i)
		i++
	}

	fmt.Println("------- For loop with range -------")

	for i := range 3 {
		fmt.Println("Range: ", i)
	}

	fmt.Println("------- For loop with the next iteration -------")

	sum = 0
	for n := range 6 {
		// Skip the odd numbers
		if n%2 != 0 {
			continue
		}

		sum += n

		fmt.Println("Next Iteration: ", n)
	}

	fmt.Println("Total Sum: ", sum)

	fmt.Println("------- For loop with range and value -------")

	// For-each range loop
	// The range form of the for loop iterates over a slice or map.
	nums := []int{2, 3, 4}
	sum = 0
	for i, num := range nums {
		sum += num
		fmt.Println("Index: ", i, " Num: ", num, " Sum: ", sum)
	}
}
