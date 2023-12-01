package main

import "fmt"

type Personable interface {
	printAge()
	changeAge(newAge int)
}

type Person struct {
	name string
	age  int
}

func (p *Person) printAge() {
	fmt.Printf("Your age is: %d\n", p.age)
}

func (p *Person) changeAge(newAge int) {
	p.age = newAge
}

func printer(text string) {
	fmt.Printf("%s", text)
}

func main() {
	felipe := Personable(&Person{})
	felipe.changeAge(30)
	felipe.printAge()
}
