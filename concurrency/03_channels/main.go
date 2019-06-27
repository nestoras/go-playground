package main

import (
	"fmt"
)

func routine(done chan bool) {
	fmt.Println("Go routine")
	// Send true to channel done.
	done <- true
}

func main() {
	done := make(chan bool)
	go routine(done)
	//Receive from channel done
	<-done
	fmt.Println("main function")
}
