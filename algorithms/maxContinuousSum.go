package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var toInt = strconv.Atoi

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	return str
}

func main() {
	n, _ := toInt(input())
	pref := make([]int, n+1)

	for i := 0; i < n; i++ {
		value, _ := toInt(input())
		pref[i+1] = pref[i] + value
	}

	maxSum := 0
	minSum := math.MaxInt
	result := [2]int{0, 0}
	for i := 1; i < n+1; i++ {
		sum := pref[i]

		if sum > maxSum {
			maxSum = sum
			result[1] = i
		}
		if sum < minSum {
			minSum = sum
			result[0] = i + 1
		}
	}
	fmt.Println(result[0], result[1])
}

//5
//-1
//-4
//3
//-2
//2
