package main

import (
	"fmt"
	ac "bank/account"
	cl "bank/client"
)

func main() {
	account := ac.CheckingAccount{}
	account.Person = cl.Person{Name: "Fred"}
	account.Balance = 500

	person2 := cl.Person{Name: "Joly"}
	account2 := ac.CheckingAccount{Person: person2, Balance: 500}

	fmt.Println(account.Withdraw(600))
	fmt.Println(account.Deposit(200))
	fmt.Println(account.Withdraw(600))
	fmt.Println(account.Transfer(100, &account2))
}

