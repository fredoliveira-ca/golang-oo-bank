package main

import (
	"fmt"
)

// CheckingAccount is a representation of a account type
type CheckingAccount struct {
	name 			string
	branch 			int
	account 		int
	accountDigit 	int
	balance 		float64
}

func main() {
	account := CheckingAccount{}
	account.name = "Fred"
	account.balance = 500

	fmt.Println(account.Withdraw(600))
	fmt.Println(account.Deposit(200))
	fmt.Println(account.Withdraw(600))
}

// Withdraw is a function for withdraw money from the balance
func (account *CheckingAccount) Withdraw(value float64) (string, string, float64) {
	isValidWithdraw := value > 0 && value <= account.balance
	if(isValidWithdraw) {
		account.balance -= value
		return "Success!", "Your balance was updated.", account.balance
	} else {
		return "Failure", "No sufficient balance.", account.balance
	}
}

// Deposit is a function for deposit money into the balance
func (account *CheckingAccount) Deposit(value float64) (string, string, float64) {
	if value > 0 {
		account.balance += value
		return "Success!", "Your balance was updated.", account.balance
	} else {
		return "Failure", "No sufficient balance.", account.balance
	}
}