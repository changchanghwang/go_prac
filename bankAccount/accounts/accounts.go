package accounts

import (
	"errors"
	"fmt"
)

// Acocunt struct
type Account struct {
	owner string
	balance int
}

var errNoMoney = errors.New("Cant't withdraw")

// NewAccount create Account
func NewAccount(owner string) *Account{
	account := Account{owner:owner, balance:0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) { // func (receiver) funcName(args)
	a.balance += amount
}

// get Balance
func (a Account) GetBalance() int{
	return a.balance
}

// Withdraw form your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil // null같은 애
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) GetOwner() string{
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.GetOwner(), "'s account.\nBalance: ",a.GetBalance())
}