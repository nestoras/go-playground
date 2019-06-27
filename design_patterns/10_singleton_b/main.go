package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

type Connection struct {
	conn string
}

var value *Connection

func Connect() *Connection {
	once.Do(func() {
		value = &Connection{conn: "Driver"}
	})

	return value
}

func main() {
	go func() {
		time.Sleep(time.Millisecond * 600)
		fmt.Println(*Connect())
	}()

	for i := 0; i < 10; i++ {

		go func(ix int) {
			time.Sleep(time.Millisecond * 60)
			fmt.Println(ix, " = ", Connect().conn)
		}(i)
	}

	fmt.Scanln()
}
