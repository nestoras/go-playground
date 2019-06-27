package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	c := make(chan bool)

	for i := 0; i < 20; i += 1 {
		m.Lock()
		go func(i int) {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println(i)
			m.Unlock()

			c <- true
		}(i)
	}
	for i := 0; i < 20; i++ {
		<-c
	}
}
