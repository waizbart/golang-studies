package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	fmt.Println("=-=-=-=-= Switch and case in golang -=-=-=-=-=")

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Dice value is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
		fallthrough
	case 2:
		fmt.Println("You can move 2 spot")
		fallthrough
	case 3:
		fmt.Println("You can move to 3 spot")
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("Dice value don't exist")
	}
}