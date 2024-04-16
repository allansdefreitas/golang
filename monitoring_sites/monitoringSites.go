package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const ExitCode = 0
const ErrorCode = -1
const TimesToTest = 5
const Delay = 2
const LogFileName = "log.txt"

func main() {

	showPresentation()

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
			showLogs()
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
			fmt.Println("Testing site ", i+1, ":", url)
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

		// fmt.Println(reflect.TypeOf(err))
		// fmt.Println(err)
		// message := siteURL + " : some problem occurred to access: " + err
		// fmt.Println(err.Error())
		errStr := string(err.Error())
		currentTime := getCurrentTimeInString()
		message := currentTime + " - " + siteURL + " : some problem occurred to access it: " + errStr

		writeAtLog(message)
	} else {

		if response.StatusCode == 200 {
			fmt.Println("The adress '", siteURL, "' was successfully loaded")
			saveLog(siteURL, true)
		} else {
			fmt.Println("The adress '", siteURL, "' is unavailable")
			saveLog(siteURL, false)
		}
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
	for {

		line, err := reader.ReadString('\n') // \n represents the byte that breaks line
		line = strings.TrimSpace(line)

		fmt.Println(line)
		sites = append(sites, line)

		if err == io.EOF { // end of line
			break
		}

	}

	arquivo.Close() // It is necessary close the file after use it to free it to other processes

	// fmt.Println(sites)
	return sites
}

func writeAtLog(message string) {

	file, err := os.OpenFile(LogFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	file.WriteString(message)
	file.WriteString("\n")
	file.Close()

}

func saveLog(siteURL string, status bool) {

	// O_RDWR: open the file read-write.
	// O_CREATE: Create a new file if not exists
	// O_APPEND append data to the file when writing.

	// permission 0666 etc
	file, err := os.OpenFile(LogFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	currentTime := getCurrentTimeInString()
	file.WriteString(currentTime + " - " + siteURL + " - online: " + strconv.FormatBool(status))
	file.WriteString("\n")
	file.Close()

}

func getCurrentTimeInString() string {

	layout := "02/01/2006 15:04:05"

	return time.Now().Format(layout)
}

func showLogs() {

	//Using ReadFile it is not necessary close file, because it already does it
	// file, err := ioutil.ReadFile(LogFileName)
	file, err := os.ReadFile(LogFileName)

	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	fmt.Println(string(file))

}
