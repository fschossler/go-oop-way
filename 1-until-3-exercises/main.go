package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

type Employee struct {
	Person
	EmployeeID int
	Salary     float64
}

func (p Person) PersonDetails() string {
	personDetails := fmt.Sprintf("Person name: %s; Person Age: %s; Person Address %s", p.Name, strconv.Itoa(p.Age), p.Address)
	return personDetails
}

func main() {
	felipe := Person{Name: "Felipe", Age: 27, Address: "Avenida Brasil, 77"}
	fmt.Println(felipe.PersonDetails())
	worker := Employee{
		Person:     Person{Name: "Joao", Age: 29, Address: "Avenida Pipoco, 28"},
		EmployeeID: 90,
		Salary:     2300.98,
	}

	fmt.Println(worker)
}
