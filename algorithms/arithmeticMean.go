package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	yearsList := strings.Split(input(), " ")

	sum := 0
	for i := 0; i < len(yearsList); i++ {
		year, _ := strconv.Atoi(yearsList[i])
		sum += year
	}

	fmt.Print(sum / len(yearsList))
}

func mainWork() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	yearsList := strings.Split(scanner.Text(), " ")

	sum := 0
	for i := 0; i < len(yearsList); i++ {
		year, _ := strconv.Atoi(yearsList[i])
		sum += year
	}

	fmt.Print(sum / len(yearsList))
}
