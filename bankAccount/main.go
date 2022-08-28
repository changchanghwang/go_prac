package main

import (
	"example/hello/bankAccount/accounts"
	"fmt"
)

func main(){
	account := accounts.NewAccount("arthur")
	account.Deposit(10000)
	fmt.Println(account.GetBalance())
}