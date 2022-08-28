package main

import (
	"example/hello/bankAccount/accounts"
	"fmt"
)

func main(){
	account := accounts.NewAccount("arthur")
	fmt.Println(account)
}