package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := sum(3, 5)

	fmt.Println("Result of sum is: ", result)

	multiplyResult, version := dot(1, 3, 5)

	fmt.Println("The multiplication of values is: ", multiplyResult)
	fmt.Println("Version: ", version)
}

func greeter() {
	fmt.Println("Hello from greeter")
}

func sum(a int, b int) int {
	return a + b
}

func dot(values ...int) (int, string) {
	total := 1

	for _, value := range values {
		total *= value
	}

	return total, "1.0"
}
