package account

import (
	"demo/json/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type AccountsVault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccountsVault() *AccountsVault {
	file, err := files.ReadFile("accounts.json")
	if err != nil {
		return &AccountsVault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault AccountsVault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return &AccountsVault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	return &vault
}

func (vault *AccountsVault) Save() {
	// Save changes
	file, err := vault.ToJSON()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	files.WriteFile("accounts.json", file)
}

func (vault *AccountsVault) ToJSON() ([]byte, error) {
	bytes, err := json.MarshalIndent(vault, "", "	")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	return bytes, nil
}

func (vault *AccountsVault) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.UpdatedAt = time.Now()

	vault.Save()
}

func (vault *AccountsVault) FindAccountByTag(tag string) []Account {
	var accounts []Account

	for _, account := range vault.Accounts {
		if strings.Contains(account.Tag, tag) {
			accounts = append(accounts, account)
		}
	}

	return accounts
}

func (vault *AccountsVault) DeleteAccountByTag(tag string) bool {
	isDeleted := false
	var newAccounts []Account
	for _, account := range vault.Accounts {
		if strings.Contains(account.Tag, tag) {
			isDeleted = true
		} else {
			newAccounts = append(newAccounts, account)
		}
	}

	vault.Accounts = newAccounts
	vault.Save()
	return isDeleted
}
