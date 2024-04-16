package main

import "fmt"

type BankAccount struct {
	holderName     string
	bankBranch     int
	accountNumber  int
	accountBalance float64
}

func main() {

	// "instantiate"
	bankAccountJohn := BankAccount{
		holderName:     "John",
		accountBalance: 8245.64,
	}

	bankAccountApolo := BankAccount{"Apolo", 123, 745216, 20554.10}

	var bankAccountPeter *BankAccount // It is necessary use '*'
	bankAccountPeter = new(BankAccount)

	bankAccountPeter.holderName = "Peter"
	bankAccountPeter.accountBalance = 22500.57

	fmt.Println(bankAccountJohn)
	fmt.Println(bankAccountApolo)
	fmt.Println("Peter: ", bankAccountPeter)  // the memory address
	fmt.Println("Peter: ", *bankAccountPeter) // the value itself
}

// Ctrl + J = show/hide terminal
// Ctrl + L (at terminal) = Clean terminal

