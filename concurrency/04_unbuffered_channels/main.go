package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Unbuffered Channel of strings.
	c := make(chan string)

	go msg("Hello!", c)

	for i := 0; i < 5; i++ {
		// Read From Channel - Blocking.
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
	}

	fmt.Println("This is the end!")
}

func msg(msg string, c chan string) {
	for i := 0; ; i++ {
		// Write to Channel.
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.

		// The write does not return until the read from main is complete.

		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}