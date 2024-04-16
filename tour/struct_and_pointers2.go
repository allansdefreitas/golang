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
	// bankAccountJohn := BankAccount{
	// 	holderName:     "John",
	// 	accountBalance: 8245.64,
	// }

	apoloBankAccount := BankAccount{"Apolo", 123, 745216, 20554.10}
	apoloBankAccount2 := BankAccount{"Apolo", 123, 745216, 20554.10}

	mariahBankAccount := BankAccount{"Mariah", 123, 745216, 20554.10}
	mariahBankAccount2 := BankAccount{"Mariah", 123, 745216, 20554.10}

	// Using pointers
	var deborahBankAccount *BankAccount // It is necessary use '*'
	deborahBankAccount = new(BankAccount)
	deborahBankAccount.holderName = "Deborah"
	deborahBankAccount.accountBalance = 20500.57

	var deborahBankAccount2 *BankAccount // It is necessary use '*'
	deborahBankAccount2 = new(BankAccount)
	deborahBankAccount2.holderName = "Deborah"
	deborahBankAccount2.accountBalance = 20500.57

	// fmt.Println(bankAccountJohn)
	// fmt.Println(bankAccountApolo)
	// fmt.Println("Peter: ", bankAccountPeter)  // the memory address
	// fmt.Println("Peter: ", *bankAccountPeter) // the value itself

	fmt.Println("Apolo")
	fmt.Println(apoloBankAccount == apoloBankAccount2)
	apoloBankAccount2.accountNumber = 1234
	fmt.Println(apoloBankAccount == apoloBankAccount2)

	fmt.Println("Mariah")
	fmt.Println(mariahBankAccount == mariahBankAccount2)
	mariahBankAccount.holderName = "Mariah of Magdalene"
	fmt.Println(mariahBankAccount == mariahBankAccount2)

	fmt.Println("Deborah")
	fmt.Println(deborahBankAccount == deborahBankAccount2)
	fmt.Println(deborahBankAccount, deborahBankAccount2)
	fmt.Println(&deborahBankAccount, &deborahBankAccount2)
	fmt.Println(deborahBankAccount == deborahBankAccount2)
	// points to the content ITSELF
	fmt.Println(*deborahBankAccount == *deborahBankAccount2)
}

// Ctrl + J = show/hide terminal
// Ctrl + L (at terminal) = Clean terminal

