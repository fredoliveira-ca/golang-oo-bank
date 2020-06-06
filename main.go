package main

import (
	"fmt"
	ac "bank/account"
	cl "bank/client"
)

func main() {
	account := ac.CheckingAccount{}
	account.Person = cl.Person{Name: "Fred"}
	account.Deposit(500)

	person2 := cl.Person{Name: "Joly"}
	account2 := ac.CheckingAccount{Person: person2}
	account2.Deposit(500)
	fmt.Println("My balance is", account.GetBalance())

	fmt.Println(account.Withdraw(600))
	fmt.Println(account.Deposit(200))
	fmt.Println(account.Withdraw(600))
	fmt.Println(account.Transfer(100, &account2))

	savingAccount := ac.SavingAccount{}
	savingAccount.Person = cl.Person{Name: "Fred"}
	savingAccount.Deposit(500)
	payBil(&savingAccount, 100)
	fmt.Println(savingAccount.GetBalance())
}

type accountInterface interface {
	Withdraw(value float64) (string, string, float64)
}

func payBil(account accountInterface, value float64) {
	account.Withdraw(value)
}
