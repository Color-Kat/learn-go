package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
	result := make([][2]int, n)

	for i := 0; i < n; i++ {
		pairStr := strings.Fields(input())
		a, _ := toInt(pairStr[0])
		b, _ := toInt(pairStr[1])
		result[i] = [2]int{a, b}
	}

	println(result)
}
