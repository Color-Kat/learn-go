package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var ToInt = strconv.Atoi

func Prompt() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func mergeTest() {
	a := []int{1, 2, 5, 6, 84}
	b := []int{2, 3, 6, 7, 40, 85}
	n := len(a)
	m := len(b)
	c := make([]int, n+m)

	l, r := 0, 0
	for l < n || r < m {
		if l < n && r < m {
			if a[l] < b[r] {
				c[l+r] = a[l]
				l++
			} else {
				c[l+r] = b[r]
				r++
			}
		} else {
			for l < n {
				c[l+r] = a[l]
				l++
			}

			for r < m {
				c[l+r] = b[r]
				r++
			}
		}
	}

	fmt.Println(c)
}

//func squareAndSort(array []int) []int {
//
//	squares := make([]int, len(array))
//
//	l, r := 0, len(array) - 1
//	index := len(array) - 1
//
//	for r >= l {
//		pow1, pow2 := array[l] * array[l], array[r] * array[r]
//
//		if pow1 >= pow2 {
//			squares[ind] = pow1
//
//		}
//	}
//
//
//	return array
//}

func maxIncreasing(nums []int) int {
	_, r := 0, 0

	maxIncrease := 0
	currentIncrease := 0

	for r < len(nums)-1 {
		if nums[r+1] > nums[r] {
			currentIncrease++
			r++

			if currentIncrease > maxIncrease {
				maxIncrease = currentIncrease
			}
		} else {
			currentIncrease = 0
			r++
		}
	}

	return maxIncrease + 1
}

func main() {

	//mergeTest()
	fmt.Println(maxIncreasing([]int{3, 1, 3, 8, 3, 1, 8, 2}))

	//n, _ := ToInt(Prompt())
	//m, _ := ToInt(Prompt())

	//for i := 0; i < n; i++ {
	//	a[i], _ = ToInt(Prompt())
	//	a[i], _ = ToInt(Prompt())
	//}
	//
	//for i := 0; i < m; i++ {
	//	b[i], _ = ToInt(Prompt())
	//}

}
