package account

import (
	"crypto/md5"
	"demo/json/files"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (a *Account) setPassword(password string) {
	hash := md5.Sum([]byte(password))
	a.Password = hex.EncodeToString(hash[:])
}

func (a *Account) ToJSON() ([]byte, error) {
	bytes, err := json.MarshalIndent(a, "", "	")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	return bytes, nil
}

func NewAccount(username, password, bio string) *Account {
	a := &Account{
		Username: username,
		Bio:      bio,
	}
	a.setPassword(password)

	return a
}

func CreateAccount() {
	myAccount := NewAccount("root", "password", "Hello, World!")
	jsonBytes, err := myAccount.ToJSON()
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
