package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Printf("Length of channel c is %v and capacity of channel c is %v", len(c), cap(c))
	// what is going here ?
//	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
//	fmt.Println(<-c)
	fmt.Printf("Length of channel c is %v and capacity of channel c is %v", len(c), cap(c))
}
