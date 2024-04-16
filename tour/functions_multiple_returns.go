package main

import "fmt"

type BankAccount struct {
	holderName     string
	bankBranch     int
	accountNumber  int
	accountBalance float64
}

func (a *BankAccount) withDrawn(value float64) (string, float64) {
	canWithdrawn := value <= a.accountBalance && value > 0

	if canWithdrawn {
		a.accountBalance -= value
		return "Successfull withdrawal", a.accountBalance
	} else {
		return "Insuficient funds", a.accountBalance
	}
}

func (a *BankAccount) deposit(value float64) (string, float64) {

	canDeposit := value > 0

	if canDeposit {
		a.accountBalance += value
		return "Successful deposit", a.accountBalance
	} else {
		return "Deposit error: Deposit value less than zero", a.accountBalance
	}
}

func main() {

	joeAccount := BankAccount{}
	joeAccount.holderName = "Joe"
	joeAccount.accountBalance = 500

	fmt.Println(joeAccount)

	fmt.Println(joeAccount.deposit(500))
	status, value := joeAccount.withDrawn(50)
	fmt.Println(status, value)
	fmt.Println(joeAccount.deposit(-500))

}

