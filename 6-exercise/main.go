package main

import "fmt"

type Calculator struct {
	number1 float32
	number2 float32
}

func (c *Calculator) addition(num1, num2 float32) float32 {
	c.number1 = num1
	c.number2 = num2
	return c.number1 + c.number2
}

func (c *Calculator) subtraction(num1, num2 float32) float32 {
	c.number1 = num1
	c.number2 = num2
	return c.number1 - c.number2
}

func (c *Calculator) multiplication(num1, num2 float32) float32 {
	c.number1 = num1
	c.number2 = num2
	return c.number1 * c.number2
}

func (c *Calculator) division(num1, num2 float32) float32 {
	c.number1 = num1
	c.number2 = num2
	return c.number1 / c.number2
}

func main() {
	sum := Calculator{number1: 10, number2: 30}
	fmt.Println(sum.addition(sum.number1, sum.number2))
}
