package main

import (
	"fmt"
	"time"
)

/*

The markers are in sequence:

ano 6 | 2006
mês 1 | 01
dia 2 | 02
hora 3 | 15
minuto 4 | 04
segundo 5 | 05

*/

func main() {

	date := time.Now()
	fmt.Println(date.Format("2 de January de 2006"))
	fmt.Println(date.Format("02/01/2006"))
	fmt.Println(date.Format("02/Jan/2006"))

}
