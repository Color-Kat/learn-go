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

func main() {
	file, _ := os.Open("INPUT.TXT")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	n, _ := ToInt(scanner.Text())

	pref := make([]int, n+1)

	for i := 0; i < n; i++ {
		scanner.Scan()
		k, _ := ToInt(scanner.Text())
		pref[i+1] = pref[i] + k
	}

	scanner.Scan()
	q, _ := ToInt(scanner.Text())

	result := make([]string, q)

	for i := 0; i < q; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		a, _ := ToInt(parts[0])
		b, _ := ToInt(parts[1])
		result[i] = strconv.Itoa(pref[b] - pref[a-1])
	}

	for i := range result {
		fmt.Println(result[i])
	}

	//fmt.Println(strings.Join(result, "-"))
	//ioutil.WriteFile("OUTPUT.TXT", []byte(
	//	strings.Join(result, "\n"),
	//), 0644)
}
