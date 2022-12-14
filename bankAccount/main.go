package main

import (
	"example/hello/bankAccount/accounts"
	"fmt"
)

func main(){
	account := accounts.NewAccount("arthur")
	account.Deposit(10000)
	err := account.Withdraw(4000)
	if err != nil {
		// log.Fatalln(err) // error와 함께 종료시킴
		fmt.Println(err)
	}
	
	account.ChangeOwner("hwang")
	fmt.Println(account)
}