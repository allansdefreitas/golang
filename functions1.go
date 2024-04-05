package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func greet(name, familyName string){
	fmt.Printf("Hi, %s %s!", name, familyName)	
}

func main() {
	fmt.Println(add(42, 13))
	greet("John", "Baptist")
}

