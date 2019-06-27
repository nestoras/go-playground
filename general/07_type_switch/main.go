package main

import (
	"fmt"
)

type Dog struct {
	Name string
}

type Person struct {
	Firstname string
	Surname string
}

type Cat struct {
	Name string
}

func describe(i interface{}) {
	switch i.(type) {
	case Dog:
		fmt.Println("Hello dog ", i.(Dog).Name)
	case Person:
		fmt.Println("Hello", i.(Person).Firstname, i.(Person).Surname)
	default:
		fmt.Println("Hello unknown", i)
	}
}


func main(){
	describe(Dog{Name:"Azor"})
	describe(Person{Firstname:"Nestoras", Surname:"Stefanou"})
	describe(Cat{Name:"Booki"})

}