package main

import "fmt"

type User struct {
	name  string
	email string
	age   int
}

func main() {
	fmt.Println("Structs in go")

	new_user := User{"gui", "gui@gmail.com", 74}

	fmt.Println("User: ", new_user)
	fmt.Printf("Formatted user: %+v\n", new_user)
}