package main

import "fmt"

type Calculator struct {
	strategy Strategy
}

func (c *Calculator) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Calculator) Calculate(a, b int) int {
	return c.strategy.Execute(a, b)
}

type Strategy interface {
	Execute(int, int) int
}

type Add struct{}

func (a *Add) Execute(value1, value2 int) int {
	return value1 + value2
}

type Subtract struct{}

func (s *Subtract) Execute(value1, value2 int) int {
	return value1 - value2
}

func main() {
	subtract := Subtract{}
	add := Add{}
	calculator := Calculator{}
	calculator.SetStrategy(&add)
	result := calculator.Calculate(3, 4)
	fmt.Println(result)
	calculator.SetStrategy(&subtract)
	result = calculator.Calculate(3, 4)
	fmt.Println(result)
}
