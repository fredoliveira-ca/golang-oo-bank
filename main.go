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

	account2 := CheckingAccount{name: "Fred", balance: 500}

	fmt.Println(account.Withdraw(600))
	fmt.Println(account.Deposit(200))
	fmt.Println(account.Withdraw(600))
	fmt.Println(account.Transfer(100, &account2))
}

// Withdraw is a function to withdraw money from the balance
func (account *CheckingAccount) Withdraw(value float64) (string, string, float64) {
	isValidWithdraw := value > 0 && value <= account.balance
	if(isValidWithdraw) {
		account.balance -= value
		return "Success Withdraw!", "Your balance was updated. Balance", account.balance
	}
	
	return "Failure Withdraw!", "No sufficient balance. Balance", account.balance
}

// Deposit is a function to deposit money into the balance
func (account *CheckingAccount) Deposit(value float64) (string, string, float64) {
	if value > 0 {
		account.balance += value
		return "Success Deposit!", "Your balance was updated. Balance", account.balance
	} 

	return "Failure Deposit!", "No sufficient balance. Balance", account.balance
}

// Transfer is a function to tranfer money between accounts
func (account *CheckingAccount) Transfer(value float64, targetAccount *CheckingAccount) (string, string, float64) {
	if account.balance >= value && value > 0 {
		account.balance -= value
		targetAccount.Deposit(value)
		return "Success Transfer!", "Your balance was updated. Balance", account.balance
	} 
	
	return "Failure Transfer!", "No sufficient balance. Balance", account.balance
}