package main

import "fmt"

type Vehicle struct{}

func (v Vehicle) Start() {
	fmt.Print("Start the engine")
}

type Car struct {
	Vehicle
}

//This is not inheritance this is embedding
func main() {
	c := Car{Vehicle{}}
	c.Start()
}
