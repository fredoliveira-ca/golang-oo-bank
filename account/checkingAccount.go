package account

import "bank/client"

// CheckingAccount is a representation of a account type
type CheckingAccount struct {
	Person 			client.Person
	Branch 			int
	Account 		int
	AccountDigit 	int
	balance 		float64
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

// GetBalance is a function to fetch the account balance
func (account *CheckingAccount) GetBalance() float64 {
	return account.balance
}