package main

import (
	"demo/json/account"
	"demo/json/utils"
	"fmt"
)

func main() {
	fmt.Println("__Accounts manager__")

	accountsVault := account.NewAccountsVault()

Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			CreateAccount(accountsVault)
		case 2:
			FindAccount(accountsVault)
		case 3:
			DeleteAccount(accountsVault)
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

func CreateAccount(accountsVault *account.AccountsVault) {
	//myAccount := NewAccount("root", "password", "Hello, World!")

	login := utils.PromptData("Enter login")
	password := utils.PromptData("Enter password")
	tag := utils.PromptData("Enter tag")
	myAccount := account.NewAccount(login, password, tag)

	accountsVault.AddAccount(*myAccount)
}

func FindAccount(accountsVault *account.AccountsVault) {
	tag := utils.PromptData("Enter tag")
	accounts := accountsVault.FindAccountByTag(tag)

	if len(accounts) == 0 {
		fmt.Println("Accounts not found")
	}

	for _, accountItem := range accounts {
		accountItem.Output()
	}
}

func DeleteAccount(accountsVault *account.AccountsVault) {
	tag := utils.PromptData("Enter tag")
	isDeleted := accountsVault.DeleteAccountByTag(tag)

	if isDeleted {
		fmt.Println("Account deleted")
	} else {
		fmt.Println("Account not found")
	}
}
