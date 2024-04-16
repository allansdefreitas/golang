package main

import (
	"fmt"
)

func main() {

	name := "Allan"

	version := 1.2

	fmt.Println("Hello,", name)
	fmt.Println("Version: ", version)

	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit program")

	var comand int
	fmt.Scan(&comand)
	//fmt.Scanf("%d", &comand) // the var type is inferred, so we don't need %d

	fmt.Println("The chosen comand was", comand)
	fmt.Println("Adress of comand is", &comand)

}

// Use the Ctrl + ` to show and hide terminal on vscode
