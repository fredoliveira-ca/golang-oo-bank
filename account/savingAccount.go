package account

import "bank/client"

// SavingAccount is a representation of a account type
type SavingAccount struct {
	Person 							client.Person
	Branch, Account, AccountDigit 	int
	balance 						float64
}

// Withdraw is a function to withdraw money from the balance
func (account *SavingAccount) Withdraw(value float64) (string, string, float64) {
	isValidWithdraw := value > 0 && value <= account.balance
	if(isValidWithdraw) {
		account.balance -= value
		return "Success Withdraw!", "Your balance was updated. Balance", account.balance
	}
	
	return "Failure Withdraw!", "No sufficient balance. Balance", account.balance
}

// Deposit is a function to deposit money into the balance
func (account *SavingAccount) Deposit(value float64) (string, string, float64) {
	if value > 0 {
		account.balance += value
		return "Success Deposit!", "Your balance was updated. Balance", account.balance
	} 

	return "Failure Deposit!", "No sufficient balance. Balance", account.balance
}

// GetBalance is a function to fetch the account balance
func (account *SavingAccount) GetBalance() float64 {
	return account.balance
}