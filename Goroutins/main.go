package main

import (
	"fmt"
)

func sum(numbers []int, channel chan int) {
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	channel <- sum
}

func main() {
	numbers := []int{1, 5, 7, 3, 1, -4, 6, -8, 38, 284, 58, 384, 91, 49, -499, 3, 4, 67, 1, 54, 8, 3, 1, 783, 28, 98, 91, 12, 3, 1, -100}

	channel := make(chan int)
	go sum(numbers[:len(numbers)/2], channel)
	go sum(numbers[len(numbers)/2:], channel)

	firstSum, secondSum := <-channel, <-channel

	fmt.Println("Sum", firstSum+secondSum)
}
