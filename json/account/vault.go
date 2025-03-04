package account

import (
	"demo/json/files"
	"encoding/json"
	"fmt"
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

func (vault *AccountsVault) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.UpdatedAt = time.Now()

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
