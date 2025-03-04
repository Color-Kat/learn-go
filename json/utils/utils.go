package utils

import "fmt"

func PromptData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

func remove[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
