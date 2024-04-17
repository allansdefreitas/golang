package main

import (
	"fmt"
)

func main(){

}

// Instead of ignoring error
func converte(s string) int {
    n, _ := strconv.Atoi(s)
    return n
}

// You can "trow" it
func converte(s string) (int, error) {
    n, err := strconv.Atoi(s)
    return n, err
}

// Or handle it
func converte(s string) (int) {
    n, err := strconv.Atoi(s)
    if err != nil {
        n = 0
        //log error
    } 
    return n
}
