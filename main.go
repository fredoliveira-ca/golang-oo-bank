package main

import (
	"fmt"
	ac "bank/account"
)

func main() {
	account := ac.CheckingAccount{}
	account.Name = "Fred"
	account.Balance = 500

	account2 := ac.CheckingAccount{Name: "Fred", Balance: 500}

	fmt.Println(account.Withdraw(600))
	fmt.Println(account.Deposit(200))
	fmt.Println(account.Withdraw(600))
	fmt.Println(account.Transfer(100, &account2))
}

