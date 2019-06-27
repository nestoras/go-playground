package main

import "fmt"

func greet(c chan string) {
	<-c
	<-c
}

func main() {
	fmt.Println("Starting....")

	c := make(chan string, 1)

	go greet(c)
	c <- "Nestoras"

	close(c)

	c <- "George"
	fmt.Println("Executed")
}
