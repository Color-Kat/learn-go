package utils

import "fmt"

func PromptData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
