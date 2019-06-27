package main

import (
	"fmt"
	"runtime"
	"time"
)

func routine() {
	fmt.Println("Go routine")
}

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	go routine()
	time.Sleep(2 * time.Second)
	fmt.Println("main function")
}
