package account

import (
	"encoding/json"
	"fmt"
	"time"
)

type AccountsVault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccountsVault() *AccountsVault {
	return &AccountsVault{
		Accounts:  []Account{},
		UpdatedAt: time.Now(),
	}
}

func (vault *AccountsVault) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.UpdatedAt = time.Now()
}

func (vault *AccountsVault) ToJSON() ([]byte, error) {
	bytes, err := json.MarshalIndent(vault, "", "	")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	return bytes, nil
}
