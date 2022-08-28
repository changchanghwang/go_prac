package accounts

import "errors"

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