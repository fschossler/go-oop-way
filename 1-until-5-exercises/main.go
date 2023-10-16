package main

import (
	"fmt"
	"strconv"
)

type Drawable interface {
	Draw() string
}

type Person struct {
	Name    string
	Age     int
	Address string
}

func (p *Person) PersonDetails() string {
	personDetails := fmt.Sprintf("Person name: %s; Person Age: %s; Person Address %s", p.Name, strconv.Itoa(p.Age), p.Address)
	return personDetails
}

func (p *Person) Draw(wordForDraw string) string {
	return "Person ########### " + wordForDraw + " ###########"
}

type Employee struct {
	Worker     Person
	EmployeeID int
	Salary     float64
}

func (e *Employee) Draw(wordForDraw string) string {
	return "Employee ########### " + wordForDraw + " ###########"
}

func main() {

	felipe := Person{Name: "Felipe", Age: 27, Address: "Avenida Brasil, 77"}
	fmt.Println(felipe.PersonDetails())

	fmt.Println(felipe.Draw("OK"))

	joao := Employee{
		Worker:     Person{Name: "Joao", Age: 29, Address: "Avenida Pipoco, 28"},
		EmployeeID: 90,
		Salary:     2300.98,
	}

	fmt.Println(joao)
	fmt.Println(joao.Draw("IEEEEEEE"))
}
