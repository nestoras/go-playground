package main

import (
	"fmt"
	"time"
)

func job1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "job1 done"
}
func job2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "job2 done"
}

func main() {
	jobChan1 := make(chan string)
	jobChan2 := make(chan string)
	go job1(jobChan1)
	go job2(jobChan2)
	select {
	case s1 := <-jobChan1:
		fmt.Println(s1)
	case s2 := <-jobChan2:
		fmt.Println(s2)
	}
}
