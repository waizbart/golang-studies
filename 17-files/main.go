package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "File that need go into a file = waizbart.dev"

	file, err := os.Create("./my-file.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./my-file.txt")
}

func readFile(filename string)  {
	dataByte, err := os.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside the file is: \n", string(dataByte))
}

func checkNilErr(err error)  {
	if err != nil {
		panic(err)
	}
}