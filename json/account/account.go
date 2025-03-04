package account

import (
	"crypto/md5"
	"demo/json/files"
	"demo/json/utils"
	"encoding/hex"
	"fmt"
)

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

func (a *Account) setPassword(password string) {
	hash := md5.Sum([]byte(password))
	a.Password = hex.EncodeToString(hash[:])
}

func NewAccount(username, password, bio string) *Account {
	a := &Account{
		Login: username,
		Url:   bio,
	}
	a.setPassword(password)

	return a
}

func CreateAccount() {
	//myAccount := NewAccount("root", "password", "Hello, World!")

	login := utils.PromptData("Enter login")
	password := utils.PromptData("Enter password")
	url := utils.PromptData("Enter url")
	myAccount := NewAccount(login, password, url)

	accountsVault := NewAccountsVault()
	accountsVault.AddAccount(*myAccount)

	jsonBytes, err := accountsVault.ToJSON()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	err = files.WriteFile("account.json", jsonBytes)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}

func FindAccount() {
	file, err := files.ReadFile("account.json")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println(string(file))
}

func DeleteAccount() {

}
