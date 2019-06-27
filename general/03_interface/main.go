package main

import "fmt"

type Bicycle struct{}

func (m Bicycle) Start() {
	fmt.Println("Start the bicycle")
}

type Car struct{}

func (c Car) Start() {
	fmt.Println("Start the car")
}

type Vehicle struct {
	StartEngine
}

type FuelControl struct {
	StartEngine
}

func (fc FuelControl) Start() {
	fmt.Println("Fuel control")
	fc.StartEngine.Start()
}

type StartEngine interface {
	Start()
}

func main() {
	v := Vehicle{FuelControl{Car{}}}
	v.Start()

	m := Vehicle{Bicycle{}}
	m.Start()
}
