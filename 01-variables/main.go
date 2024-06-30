package main

import "fmt"

// capital T turns the variable public
const TIMEOUT int = 300

func main() {
	var username string = "waizber"
	fmt.Println(username)
	fmt.Printf("Variable username is of type: %T \n", username)

	var isLoading bool = true
	fmt.Println(isLoading)
	fmt.Printf("Variable isLoading is of type: %T \n", isLoading)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Variable smallInt is of type: %T \n", smallInt)

	var smallFloat float64 = 255.3636363463776334
	fmt.Println(smallFloat)
	fmt.Printf("Variable smallFloat is of type: %T \n", smallFloat)
	
	// default values and some aliases
	var anotherVar int 
	fmt.Println(anotherVar) // 0
	fmt.Printf("Variable anotherVar is of type: %T \n", anotherVar)

	var strVar string 
	fmt.Println(strVar) // ""
	fmt.Printf("Variable strVar is of type: %T \n", strVar)

	// implicit types
	var website = "waizbart.dev" // type string
	fmt.Println(website)

	// no var style
	numberOfDucks := 1000000
	fmt.Println(numberOfDucks)

	fmt.Println(TIMEOUT)
}