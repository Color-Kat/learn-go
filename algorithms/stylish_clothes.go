package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var ToInt = strconv.Atoi

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()

	return str
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {

	n, _ := ToInt(Input())
	partsA := strings.Fields(Input())
	a := make([]int, n)

	for i, value := range partsA {
		a[i], _ = ToInt(value)
	}

	m, _ := ToInt(Input())
	partsB := strings.Fields(Input())
	b := make([]int, m)

	for i, value := range partsB {
		b[i], _ = ToInt(value)
	}

	// Logic

	l, r := 0, 0
	bestScore := math.MaxInt64
	var bestScoreValueA, bestScoreValueB int

	for bestScore != 0 && l < n && r < m {
		currentScore := Abs(a[l] - b[r])
		if currentScore < bestScore {
			bestScore = currentScore
			bestScoreValueA = a[l]
			bestScoreValueB = b[r]
		}

		if a[l] > b[r] {
			r++
		} else {
			l++
		}
	}

	fmt.Print(bestScoreValueA, bestScoreValueB)

}
