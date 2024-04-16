package main

// Can use a nick of a package
import (
	"2_bank/accounts"
	"fmt"
)

func PayBankSlip(account verifyAccount, valueofSlip float64) {
	account.WithDrawn(valueofSlip)
}

type verifyAccount interface {
	WithDrawn(value float64) (string, float64)
}

// go mod init 2_bank
func main() {

	mikeAccount := accounts.SavingsAccount{}

	mikeAccount.Deposit(100)

	PayBankSlip(&mikeAccount, 60)

	fmt.Println(mikeAccount)

	juniorAccount := accounts.BankAccount{}
	juniorAccount.Deposit(500)

	PayBankSlip(&juniorAccount, 300)
}
