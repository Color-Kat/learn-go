package account

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Tag      string `json:"tag"`
}

func NewAccount(username, password, tag string) *Account {
	a := &Account{
		Login: username,
		Tag:   tag,
	}
	a.setPassword(password)

	return a
}

func (a *Account) setPassword(password string) {
	hash := md5.Sum([]byte(password))
	a.Password = hex.EncodeToString(hash[:])
}

func (a *Account) Output() {
	fmt.Println("--- " + a.Tag + " ---")
	fmt.Println("Login:", a.Login)
	fmt.Println("Password:", a.Password)
}
