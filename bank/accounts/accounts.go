package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (account *Account) Deposit(amount int) {
	account.balance += amount
}

// Balance of your account
func (account Account) Balance() int {
	return account.balance
}

// Withdraw x amount from your account
func (account *Account) Withdraw(amount int) error {
	if account.balance < amount {
		return errNoMoney
	}
	account.balance -= amount
	return nil
}

// ChangeOwner of the account
func (account *Account) ChangeOwner(newOwner string) {
	account.owner = newOwner
}

// Owner of the account
func (account Account) Owner() string {
	return account.owner
}

func (account Account) String() string {
	return fmt.Sprint(account.Owner(), "'s account.\nHas: ", account.Balance())
}
