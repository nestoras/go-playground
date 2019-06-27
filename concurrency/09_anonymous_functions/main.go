package main

import "fmt"

func main() {
	fmt.Println("Starting...")
	c := make(chan string)

	go func(c chan string) {
		fmt.Println("Hello " + <-c + "!")
	}(c)

	c <- "Nestoras"
}
