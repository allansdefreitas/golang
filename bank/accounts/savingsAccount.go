package accounts

import "2_bank/clients"

type SavingsAccount struct {
	Holder                               clients.Client
	BankBranch, AccountNumber, Operation int
	AccountBalance                       float64
}

func (a *SavingsAccount) WithDrawn(value float64) (string, float64) {
	canWithdrawn := value <= a.AccountBalance && value > 0

	if canWithdrawn {
		a.AccountBalance -= value
		return "Successfull withdrawal", a.AccountBalance
	} else {
		return "Insuficient funds", a.AccountBalance
	}
}

func (a *SavingsAccount) Deposit(value float64) (string, float64) {

	canDeposit := value > 0

	if canDeposit {
		a.AccountBalance += value
		return "Successful Deposit", a.AccountBalance
	} else {
		return "Deposit error: Deposit value less than zero", a.AccountBalance
	}
}

func (a *SavingsAccount) Transfer(value float64, receiver *SavingsAccount) bool {

	var canTransfer bool = value <= a.AccountBalance && value > 0

	if canTransfer {

		// a.AccountBalance -= value
		a.WithDrawn(value)
		receiver.Deposit(value)
		return true

	} else {
		return false
	}
}

func (a *SavingsAccount) GetAccountBalance() float64 {

	return a.AccountBalance
}
