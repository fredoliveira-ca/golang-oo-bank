package main

import (
	"fmt"
)

type CheckingAccount struct {
	name 			string
	branch 			int
	account 		int
	accountDigit 	int
	balance 		float64
}

func main() {
	checkingAccount := CheckingAccount{
		name: "Fred Oliveira",
		branch: 1234,
		account: 58096,
		accountDigit: 1,
		balance: 897.32}

	checkingAccount1 := CheckingAccount{
		name: "Fred Oliveira",
		balance: 897.32}
	
	checkingAccount2 := CheckingAccount{ "Fred Oliveira", 1234, 58096, 1, 897.32 }
	
	fmt.Println(checkingAccount, checkingAccount1, checkingAccount2)

	var someoneAccount *CheckingAccount
	someoneAccount = new(CheckingAccount)
	someoneAccount.name = "Jeseph"
	someoneAccount.balance = 432.12

	fmt.Println(&someoneAccount)
	fmt.Println(*someoneAccount)
	fmt.Println(someoneAccount)

}