package accounts

import "2_bank/clients"

/*
Using First character of an attribute or function in Upper case,
the visibility becomes public to another packages
*/
type BankAccount struct {
	Holder                    clients.Client
	BankBranch, AccountNumber int
	accountBalance            float64
}

func (a *BankAccount) WithDrawn(value float64) (string, float64) {
	canWithdrawn := value <= a.accountBalance && value > 0

	if canWithdrawn {
		a.accountBalance -= value
		return "Successfull withdrawal", a.accountBalance
	} else {
		return "Insuficient funds", a.accountBalance
	}
}

func (a *BankAccount) Deposit(value float64) (string, float64) {

	canDeposit := value > 0

	if canDeposit {
		a.accountBalance += value
		return "Successful Deposit", a.accountBalance
	} else {
		return "Deposit error: Deposit value less than zero", a.accountBalance
	}
}

// (a BankAccout) uses passing by value
// (a *BankAccount) uses passing by reference
func (a *BankAccount) Transfer(value float64, receiver *BankAccount) bool {

	var canTransfer bool = value <= a.accountBalance && value > 0

	if canTransfer {

		// a.accountBalance -= value
		a.WithDrawn(value)
		receiver.Deposit(value)
		return true

	} else {
		return false
	}
}

func (a *BankAccount) GetAccountBalance() float64 {

	return a.accountBalance
}
