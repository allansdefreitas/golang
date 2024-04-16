package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "Allan" // declare variables with := instead of 'var'
	age := 29
	version := 1.2

	fmt.Println("Hello,", name, "your age is ", age)
	fmt.Println("Version: ", version)

	fmt.Println("Type of name is", reflect.TypeOf(name))
	fmt.Println("Type of age is", reflect.TypeOf(age))
	fmt.Println("Type of version is", reflect.TypeOf(version))
}

// Use the Ctrl + ` to show and hide terminal on vscode
