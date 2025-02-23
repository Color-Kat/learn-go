package main

import (
	"fmt"
	"math/rand"
)

func doSomeMath(mark string, channel chan string) {
	result := fmt.Sprintln(mark, rand.Intn(10)+rand.Intn(10))
	channel <- result
}

func main() {
	channel := make(chan string)

	go doSomeMath("1", channel)
	go doSomeMath("2", channel)
	go doSomeMath("3", channel)

	res1, res2, res3 := <-channel, <-channel, <-channel

	fmt.Println("end", res1, res2, res3)
}
