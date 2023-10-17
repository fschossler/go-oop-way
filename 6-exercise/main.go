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

type AdvancedCalculator struct {
	BasicArimeticOperators Calculator
}

func (a *AdvancedCalculator) squareRoot(num1, num2 float32) float32 {
	a.BasicArimeticOperators.number1 = num1
	a.BasicArimeticOperators.number2 = num2

	for i := 0; i < exponent; i++ {
		result *= base
	}

	return result
}

func main() {
	sum := Calculator{number1: 10, number2: 30}
	fmt.Println(sum.addition(sum.number1, sum.number2))

	power := AdvancedCalculator{
		BasicArimeticOperators: Calculator{number1: 30, number2: 50},
	}
}
