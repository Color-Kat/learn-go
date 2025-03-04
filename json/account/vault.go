package account

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write(data []byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func NewAccountsVault(db Db) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}

func (vault *VaultWithDb) Save() {
	vault.UpdatedAt = time.Now()

	// Save changes only in Vault, not in VaultWithDb
	file, err := vault.Vault.ToJSON()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	vault.db.Write(file)
}

// ToJSON convert Vault (not VaultWithDb) to JSON
func (vault *Vault) ToJSON() ([]byte, error) {
	bytes, err := json.MarshalIndent(vault, "", "	")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	return bytes, nil
}

func (vault *VaultWithDb) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.Save()
}

func (vault *VaultWithDb) FindAccountByTag(tag string) []Account {
	var accounts []Account

	for _, account := range vault.Accounts {
		if strings.Contains(account.Tag, tag) {
			accounts = append(accounts, account)
		}
	}

	return accounts
}

func (vault *VaultWithDb) DeleteAccountByTag(tag string) bool {
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
