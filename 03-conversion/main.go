package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our BJJ app")
	fmt.Print("Please rate our school between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Rate: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error at conversion: ", err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating + 1)
	}
}