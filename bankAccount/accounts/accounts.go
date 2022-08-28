package accounts

// Acocunt struct
type Account struct {
	owner string
	balance int
}

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