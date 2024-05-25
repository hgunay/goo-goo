package main

import "fmt"

type employee struct {
	fullName string
	age      int
	salary   float64
}

func main() {
	employee1 := employee{"Hakan Gunay", 40, 5000}
	fmt.Println(employee1)

	employee2 := employee{
		fullName: "Hasan Gunay",
		age:      44,
		salary:   6000,
	}
	fmt.Println(employee2)

	employee3 := employee{
		fullName: "Burcu Gunay",
		age:      39,
	}
	fmt.Println(employee3)
}
