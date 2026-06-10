package main

import "fmt"

func main() {
	fmt.Println("Welcome to my pointers studies")

	var ptr1 *int
	fmt.Println("Value of pointer is: ", ptr1)

	myNumber := 19

	var ptr2 = &myNumber
    fmt.Println("Value of actual address is", ptr2)
    fmt.Println("Value of actual pointer is", *ptr2)

	*ptr2 = *ptr2 * 2

	fmt.Println("New value is: ", myNumber)
}