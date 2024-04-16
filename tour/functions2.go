package main

import "fmt"
import "strings"
import "strconv"

func swap(x, y string) (string, string) {
	return y, x
}

func citeAuthor(name, familyName string) (string){
	citing := strings.ToUpper(familyName) + ", " + name
	return citing
}

func cite(familyName string, year int) (string){
	citing := strings.ToUpper(familyName) + ", " + strconv.Itoa(year)
	return citing
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(citeAuthor("Allan", "Freitas"))
	fmt.Println(cite("Freitas", 2024))
}

