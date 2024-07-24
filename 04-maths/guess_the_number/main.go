package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func ask_guess(random_n int64, tries int) int {
	tries++

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

	return ask_guess(random_n, tries)
}

func main()  {
	fmt.Println("===== GUESS THE NUMBER I'M THINKING! =====")

	var random_n int = rand.Intn(10)

	tries := ask_guess(int64(random_n), 0)

	fmt.Println("CONGRATULATIONS, YOUR ANSWER IS RIGHT! Tries count: ", tries)

	fmt.Scanln()
}