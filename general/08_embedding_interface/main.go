package main

import "fmt"

type IAction interface {
	Walker
	Sleeper
}

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
	var a IAction = person
	var w Walker = person
	var s Sleeper = person
	fmt.Printf("dynamic type and value of interface a of static type IAction is'%T' and '%v'\n", a, a)
	fmt.Printf("dynamic type and value of interface w of static type Walker is'%T' and '%v'\n", w, w)
	fmt.Printf("dynamic type and value of interface s of static type Sleeper is'%T' and '%v'\n", s, s)
}




