package main

import (
	"fmt"
	"math/big"
	//"math/rand"
	"crypto/rand"
	// "time"
)

func main()  {
	fmt.Println("Welcome to my math studies")

	//var mynumber1 int = 2
	// var mynumber2 float64 = 7 // don't need to be 7.0
	// fmt.Println("The sum is: ", mynumber1 + int(mynumber2))

	// random number
	// rand.Seed(time.Now().UnixNano()) // deprecated
	// random_n := rand.Intn(10) + 1
	// fmt.Println("Random number: ", random_n)

	// random from crypto
	randN, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randN)
}