package main

import "fmt"

type Walker interface {
	Walk()
}

type Sleeper interface {
	Sleep()
}

type IFly interface {
	Fly()
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

//func (p Person) Fly() {
//	fmt.Println(p.Name + " flying...")
//}


func main(){
	var w Walker = Person{Name:"Nestoras"}
	_, ok1 := w.(Walker)
	fmt.Printf("implements interface Walker? %v\n", ok1)
	_, ok2 := w.(IFly)
	fmt.Printf("implements interface IFly? %v\n", ok2)
}