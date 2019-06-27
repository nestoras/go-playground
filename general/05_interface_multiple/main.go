package main

import "fmt"

type Walker interface {
	Walk()
}

type Sleeper interface {
	Sleep()
}

type Person struct {
	Name string
}

func (p Person) Walk() {
	fmt.Println(p.Name + " walking...")
}

func (p Person) Sleep() {
	fmt.Println(p.Name + " sleeping...")
}


func main(){
	person := Person{Name:"Nestoras"}
	person.Sleep()
	person.Walk()
}