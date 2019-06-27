package main

import (
	"fmt"
	"sync"
)

var number = 0

func add(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	number = number + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go add(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of number", number)
}
