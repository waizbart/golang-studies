package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const RANGE = 100

func ask_guess(random_n int64, tries int) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("your guess: ")
	read_guess, _ := reader.ReadString('\n')
	guess, _ := strconv.ParseInt(strings.TrimSpace(read_guess), 10, 64)

	if guess == random_n {
		return tries
	} else if guess < random_n {
		fmt.Println("go up...")
	} else {
		fmt.Println("go down...")
	}

	return ask_guess(random_n, tries + 1)
}

func main()  {
	fmt.Printf("===== GUESS THE NUMBER I'M THINKING! (0-%d) =====\n\n", RANGE)

	var random_n int = rand.Intn(RANGE)

	tries := ask_guess(int64(random_n), 1)

	fmt.Println("CONGRATULATIONS, YOUR ANSWER IS RIGHT! Tries count: ", tries)
}