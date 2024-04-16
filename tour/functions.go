package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	const ExitCode = 0
	const ErrorCode = -1

	name, _ := returnNameAndAge()

	fmt.Println(name)

	showPresentation()

	fmt.Print("\n")

	showMenu()

	fmt.Print("\n")
	fmt.Println("Enter you command: ")

	command := readCommand()
	fmt.Print("\n")

	switch command {

	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Showing logs...")
	case 0:
		fmt.Println("Exiting program...")
		os.Exit(ExitCode)
	default:
		fmt.Println("This option does not exist!" +
			"\nPlease select a valid option.")
		os.Exit(ErrorCode)
	}

}

func showPresentation() {

	name := "Allan"
	version := 1.2

	fmt.Println("Hello,", name)
	fmt.Println("Version: ", version)

}

func showMenu() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit program")
}

func readCommand() int {

	var command int
	fmt.Scan(&command)
	fmt.Println("The chosen command was", command)

	return command
}

func startMonitoring() {

	fmt.Println("Monitoring...")
	siteURL := "https://www.alura.com.br"

	// response, err := http.Get(siteURL)

	fmt.Println(http.Get(siteURL))

}

// Multiple return function
func returnNameAndAge() (string, int) {

	name := "Allan"
	age := 29
	return name, age
}
