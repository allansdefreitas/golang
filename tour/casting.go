package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	salario := "1000"
	aumento := "20"

	fmt.Println("O salário é", salario, "e o aumento de", aumento+"%")

	convertedSalary, err := strconv.Atoi(salario)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reflect.TypeOf(convertedSalary), convertedSalary)

	porcentageIncrease, err := strconv.Atoi(aumento)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(porcentageIncrease), porcentageIncrease)

	fmt.Println((porcentageIncrease / 100))
	newSalary := convertedSalary * (1 + (porcentageIncrease / 100))

	fmt.Println("New salary: $", newSalary)

	fmt.Println("Wrong! We need to convert to a decimal number!!!")

	porcentageIncreaseFloat, _ := strconv.ParseFloat(aumento, 64)
	convertedSalaryFloat, _ := strconv.ParseFloat(salario, 64)

	fmt.Println(porcentageIncreaseFloat, convertedSalaryFloat)

	fmt.Println(((porcentageIncreaseFloat) / 100))
	newSalaryFloat := (convertedSalaryFloat * (1 + (porcentageIncreaseFloat)/100))

	fmt.Printf("New salary: $%.2f is well calculated now!\n", newSalaryFloat)
}
