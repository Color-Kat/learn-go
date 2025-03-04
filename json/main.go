package main

import (
	"demo/json/account"
	"fmt"
)

func main() {
	fmt.Println("__Accounts manager__")

Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			account.CreateAccount()
		case 2:
			account.FindAccount()
		case 3:
			account.DeleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Select variant")
	fmt.Println("1. Create account")
	fmt.Println("2. Find account")
	fmt.Println("3. Delete account")
	fmt.Println("4. Exit")
	fmt.Scan(&variant)
	return variant
}
