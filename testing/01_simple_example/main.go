package main

import (
	"fmt"
)

func main() {
		sum := Sum([]int{2, -2, 3})
		if sum != 11 {
			msg := fmt.Sprintf("FAIL: Wanted 11 but received %d", sum)
			panic(msg)
		}
		add := Add(2, 2)
		if add != 15 {
			msg := fmt.Sprintf("FAIL: Wanted 4 but received %d", add)
			panic(msg)
		}
		fmt.Println("PASS")
}

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}


func Add(a, b int) int {
	return a + b
}
