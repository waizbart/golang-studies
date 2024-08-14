package main

import "fmt"

func main()  {
	fmt.Println("If else in golang")

	loginCount := 4

	var exceeded bool = false

	if loginCount > 10 {
		exceeded = true
	}

	if exceeded {
		fmt.Println("login attempts exceeded")
	} else {
		fmt.Println("user can login")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than ten")
	} else {
		fmt.Println("Num is not less than ten")
	}
}