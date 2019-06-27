package main

import "fmt"

func sendData(sendChannel chan <-int) {
	sendChannel <- 10
}

func main() {
	sendChannel := make(chan <-int)
	go sendData(sendChannel)
	fmt.Println(<-sendChannel)
}
