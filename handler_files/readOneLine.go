package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const ExitCode = 0
const ErrorCode = -1
const TimesToTest = 5
const Delay = 5

func main() {

	showPresentation()
	readSitesURLFromFile()

	fmt.Print("\n")

	for { // for without condition runs ad aeternum

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
	fmt.Println("")

	return command
}

func startMonitoring() {

	fmt.Println("Monitoring...")

	siteURLs := readSitesURLFromFile()

	for i := 0; i < TimesToTest; i++ {

		for i, url := range siteURLs {
			fmt.Println("Testint site ", i+1, ":", url)
			testSiteURL(url)
			fmt.Print("\n")
		}
		time.Sleep(Delay * time.Second)
		fmt.Println("")
	}
}

func testSiteURL(siteURL string) {
	response, err := http.Get(siteURL)

	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("The adress '", siteURL, "' was successfully loaded")
	} else {
		fmt.Println("The adress '", siteURL, "' is unavailable")
	}
}

func readSitesURLFromFile() []string {

	var sites []string

	arquivo, err := os.Open("sitesURL.txt")
	// arquivo, err := ioutil.ReadFile("sitesURL.txt") //deprecated
	// arquivo, err := os.ReadFile("sitesURL.txt")
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	reader := bufio.NewReader(arquivo)

	line, err := reader.ReadString('\n') // \n represents the byte that breaks line

	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	fmt.Println(line)

	return sites
}
