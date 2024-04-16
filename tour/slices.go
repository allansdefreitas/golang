package main

import (
	"fmt"
	"reflect"
)

func main() {

	namesSlice := []string{"Peter", "James", "John"}
	var namesArray [3]string

	namesArray[0] = "Peter"
	namesArray[1] = "James"
	namesArray[2] = "John"

	fmt.Println("Array: ", namesArray)
	fmt.Println("Array length:", len(namesArray))
	fmt.Println("Array capacity: ", cap(namesArray))

	fmt.Println("Slice: ", namesSlice)
	fmt.Println("Slice length:", len(namesSlice))
	fmt.Println("Slice capacity: ", cap(namesSlice))

	namesSlice = append(namesSlice, "Paul")

	fmt.Println(namesSlice)
	fmt.Println("Slice length:", len(namesSlice))
	fmt.Println("Slice capacity: ", cap(namesSlice))

	fmt.Println("\nType of Slice: ", reflect.TypeOf(namesSlice))
	fmt.Println("Type of Array: ", reflect.TypeOf(namesArray))
	// Slice is just an abstraction of an Array

}
