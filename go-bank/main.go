package main

import (
	"fmt"

	"github.com/devheesukang/learngo/go-bank/accounts"
)

func main() {
	// account := banking.Account{Owner: "heesu", Balance: 1000}
	// fmt.Println(account)
	account := accounts.NewAccount("heesu")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
