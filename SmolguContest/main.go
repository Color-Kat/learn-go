package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ToInt = strconv.Atoi

func Prompt() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func MapSlice[T any, U any](slice []T, callback func(elem T) U) []U {
	result := make([]U, len(slice))

	for i, value := range slice {
		result[i] = callback(value)
	}

	return result
}

var intArray = MapSlice[string, int]([]string{"1", "2", "4"}, func(str string) int {
	value, _ := ToInt(str)
	return value
})

// H
// m minutes
// n logs
// t = 0 - start
// add log 1 min before last log is gone
// Burn all logs

func solve() string {
	file, _ := os.Open("INPUT.TXT")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	parts1 := strings.Fields(scanner.Text())
	n, _ := ToInt(parts1[0])
	desiredBurnTime, _ := ToInt(parts1[1])

	scanner.Scan()
	logBurnTimesString := strings.Fields(scanner.Text())
	logBurnTimes := make([]int, n)

	for i := 0; i < n; i++ {
		logBurnTimes[i], _ = ToInt(logBurnTimesString[i])
		if logBurnTimes[i] > desiredBurnTime {
			return "no"
		}
	}

	currentBurnTime := 1
	for i := 0; i < n; i++ {

		nextBurnTime := currentBurnTime + logBurnTimes[i] - 1

		if nextBurnTime <= desiredBurnTime {
			currentBurnTime = nextBurnTime
		} else {
			currentBurnTime = desiredBurnTime
		}
	}

	if currentBurnTime == desiredBurnTime {
		return "yes"
	} else {
		return "no"
	}
}

func main() {

	//fmt.Println(solve())

	//file, _ := os.Open("INPUT.TXT")
	//defer file.Close()

	// Чтение файла построчно
	//scanner := bufio.NewScanner(file)
	//scanner.Scan()
	//for scanner.Scan() {
	//	side := scanner.Text()
	//}

	os.WriteFile("OUTPUT.TXT", []byte(fmt.Sprint(
		solve(),
	)), 0644)
}
