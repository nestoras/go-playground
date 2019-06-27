package main

import (
	"fmt"
)

type Vehicle struct {
	state State
}

func (v *Vehicle) Init() {
	v.state.Start()
}

func (v *Vehicle) Destroy() {
	v.state.Stop()
}

func (v *Vehicle) SetState(state State) {
	v.state = state
}

type State interface {
	Start()
	Stop()
}

type Car struct{}

func (c *Car) Start() {
	fmt.Println("A new car")
}

func (c *Car) Stop() {
	fmt.Println("Stop the car")
}

type Bicycle struct{}

func (b *Bicycle) Start() {
	fmt.Println("A new bicycle")
}

func (b *Bicycle) Stop() {
	fmt.Println("Stop the bicycle")
}

func main() {

	car := Car{}
	bicycle := Bicycle{}
	vehicle := Vehicle{&car}
	vehicle.Init()
	vehicle.Destroy()
	vehicle.SetState(&bicycle)
	vehicle.Init()
	vehicle.Destroy()
}
