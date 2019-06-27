package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(4)
	go helloWorldRepeat()
	go hello()
	go world()
	go helloWorld()
	wg.Wait()
}

func hello() {
	fmt.Println("Hello")
	wg.Done()
}

func world() {
	fmt.Println("World")
	wg.Done()
}

func helloWorld() {
	fmt.Println("Hello World!!")
	wg.Done()
}

func helloWorldRepeat() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World!!!!!!!!!")
	}
	wg.Done()
}
