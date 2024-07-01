package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome message blablabla"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter you dream:")
	

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for give us your dream description. We hope it comes real! :)")
	fmt.Println("Received dream description: ", input)
	fmt.Printf("Type of this dream is %T", input)
}