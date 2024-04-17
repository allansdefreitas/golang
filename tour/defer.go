package main

import "fmt"

func main() {
	defer fmt.Println("last last")
	defer fmt.Println("last")

	fmt.Println("first")
	fmt.Println("second")
}

