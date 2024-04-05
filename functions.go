package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int{
	return x - y
}

func mul(x float32, y float32) float32{
	return x * y	
}

func div(x float64, y float64) float64{
	return x / y	
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(10, 5))
	fmt.Println(mul(10, 2))
	fmt.Println(div(10, 3))
}

