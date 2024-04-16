package main

import (
	"fmt"
)

func main() {

	name := "Allan"

	version := 1.2

	fmt.Println("Hello,", name)
	fmt.Println("Version: ", version)

	fmt.Print("\n")

	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit program")

	fmt.Print("\n")

	fmt.Println("Enter you command: ")
	var command int
	fmt.Scan(&command)

	fmt.Println("The chosen command was", command)

	fmt.Print("\n")

	switch command {

	case 1:
		fmt.Println("Monitoring...") // we don't need break
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Exiting program...")
	default:
		fmt.Println("This option does not exist!" +
			"\nPlease select a valid option.")
	}

}

// Use the Ctrl + ` to show and hide terminal on vscode
