package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go msg("Hello!")
	//msg("Hello!")
}

func msg(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}