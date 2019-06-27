package main

import "fmt"

type Dog struct{
	Name string
}

type Car struct {
	Name  *string
	Model string
	Manufacturer string
}


func describe (i interface{}) {
	fmt.Printf("This is type of %T with value %v\n", i ,i)
}

func main(){
	dog := Dog{Name:"Azor"}
	car := Car{Model:"I8", Manufacturer:"BMW"}
	describe(car)
	describe(dog)
}