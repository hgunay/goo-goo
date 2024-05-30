package main

import "fmt"

func main() {
	// Declaring a map
	var employeeSalary map[string]int
	fmt.Println("Declaring a map\n", employeeSalary)

	// Initializing a map
	employeeSalary2 := make(map[string]int)
	fmt.Println("Initializing a map\n", employeeSalary2)

	// Initializing a map with values
	employeeSalary3 := map[string]int{
		"Hakan": 12000,
		"Burcu": 15000,
	}
	fmt.Println("Map with values\n", employeeSalary3)

	// Add
	employeeSalary3["Hasan"] = 13000
	fmt.Println("Adding an employee to map", employeeSalary3)

	// Get
	fmt.Printf("Getting Hakan's salary: %d\n", employeeSalary3["Hakan"])

	// Delete
	delete(employeeSalary3, "Hasan")

	fmt.Println("Listing employeeSalary3 map after deleting Hasan\n", employeeSalary3)
}
