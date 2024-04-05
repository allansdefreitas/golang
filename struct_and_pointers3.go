package main

import "fmt"

type BankAccount struct {
	holderName     string
	bankBranch     int
	accountNumber  int
	accountBalance float64
}

func (a *BankAccount) withDrawn(withdrawalValue float64) string {
	canWithdrawn := withdrawalValue <= a.accountBalance && withdrawalValue > 0

	if canWithdrawn {
		a.accountBalance -= withdrawalValue
		return "Successfull withdrawal"
	} else {
		return "Insuficient funds"
	}
}

func main() {

	elishaAccount := BankAccount{}
	elishaAccount.holderName = "Elisha"
	elishaAccount.accountBalance = 700

	fmt.Println(elishaAccount.withDrawn(-100))
	fmt.Println(elishaAccount.withDrawn(-100))
	fmt.Println(elishaAccount.withDrawn(-100))

	fmt.Println(elishaAccount.accountBalance)
}

