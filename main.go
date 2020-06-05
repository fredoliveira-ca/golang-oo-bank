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

}

// Withdraw is a function for withdraw money from the balance
func (account *CheckingAccount) Withdraw(value float64) string {
	isValidWithdraw := value > 0 && value <= account.balance
	if(isValidWithdraw) {
		account.balance -= value
		return "Success!" + " Your new balance is: " + fmt.Sprintf("%f", account.balance)
	} else {
		return "No sufficient balance." + " Your new balance is: " + fmt.Sprintf("%f", account.balance)
	}
}