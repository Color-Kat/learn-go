package main

import "fmt"

// fibonacci using closure
func fibonacci() func() int {
	prev := 0
	current := 1

	return func() int {
		prev, current = current, prev+current

		return current
	}
}

func main() {
	f1, f2 := fibonacci(), fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f1(), f2())
	}
}
