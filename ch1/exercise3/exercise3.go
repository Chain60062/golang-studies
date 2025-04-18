package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {

	employee1 := Employee{"Elizabeth", "Smith", 1203}

	employee2 := Employee{
		firstName: "Edward",
		lastName:  "Williams",
		id:        2146,
	}
	var employee3 Employee

	employee3.firstName = "Judy"
	employee3.lastName = "Garcia"
	employee3.id = 958

	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee3)
}
