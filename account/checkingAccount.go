package account

// CheckingAccount is a representation of a account type
type CheckingAccount struct {
	Name 			string
	Branch 			int
	Account 		int
	AccountDigit 	int
	Balance 		float64
}

// Withdraw is a function to withdraw money from the balance
func (account *CheckingAccount) Withdraw(value float64) (string, string, float64) {
	isValidWithdraw := value > 0 && value <= account.Balance
	if(isValidWithdraw) {
		account.Balance -= value
		return "Success Withdraw!", "Your balance was updated. Balance", account.Balance
	}
	
	return "Failure Withdraw!", "No sufficient balance. Balance", account.Balance
}

// Deposit is a function to deposit money into the balance
func (account *CheckingAccount) Deposit(value float64) (string, string, float64) {
	if value > 0 {
		account.Balance += value
		return "Success Deposit!", "Your balance was updated. Balance", account.Balance
	} 

	return "Failure Deposit!", "No sufficient balance. Balance", account.Balance
}

// Transfer is a function to tranfer money between accounts
func (account *CheckingAccount) Transfer(value float64, targetAccount *CheckingAccount) (string, string, float64) {
	if account.Balance >= value && value > 0 {
		account.Balance -= value
		targetAccount.Deposit(value)
		return "Success Transfer!", "Your balance was updated. Balance", account.Balance
	} 
	
	return "Failure Transfer!", "No sufficient balance. Balance", account.Balance
}