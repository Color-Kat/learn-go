package main

import "fmt"

func lastWordLength(str string) int {
	length := 0
	for i := len(str) - 1; i > 0; i-- {
		if length == 0 && str[i] == ' ' {
			continue
		}

		if str[i] != ' ' {
			length++
		} else {
			break
		}
	}

	return length
}

func increaseSliceNumber(slice []int) []int {
	// Add new 0 digit at slice start
	slice = append([]int{0}, slice...)

	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] < 9 {
			slice[i]++
			break
		} else {
			slice[i] = 0
		}
	}

	// No first digit
	if slice[0] == 0 {
		return slice[1:]
	}

	return slice
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reverse := 0
	copy := x

	for i := 1; copy*10 > 1; i *= 10 {
		reverse = reverse*10 + copy%10
		copy /= 10
	}

	return x == reverse
}

func main() {
	//// Task 1
	//fmt.Println(lastWordLength("I love typescript                 "))
	//
	//// Task 2
	//fmt.Println(increaseSliceNumber([]int{9, 9}))
	//fmt.Println(increaseSliceNumber([]int{9}))
	//fmt.Println(increaseSliceNumber([]int{9, 9, 9, 9}))

	// Task 3
	fmt.Println(isPalindrome(1221))
}

// [1,2,3,4,9]
