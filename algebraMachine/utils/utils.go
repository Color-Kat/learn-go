package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

var ToInt = strconv.Atoi

func Prompt(prompt string) string {
	fmt.Println(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	return input
}

func MapSlice[T any, U any](slice []T, callback func(elem T) U) []U {
	result := make([]U, len(slice))

	for i, value := range slice {
		result[i] = callback(value)
	}

	return result
}

func ClearConsole() {
	cmd := exec.Command("cmd", "/c", "cls") // Windows
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
