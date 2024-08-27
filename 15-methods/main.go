package main

import "fmt"

type User struct {
	id     int
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u *User) NewMail(email string) {
	u.Email = email
}

func main() {
	fmt.Println("Structs in go")

	new_user := User{1, "gui", "gui@gmail.com", 74, true}

	fmt.Println("User: ", new_user)
	fmt.Printf("Formatted user: %+v\n", new_user)
	new_user.GetStatus()
	new_user.NewMail("joao@gmail.com")

	fmt.Println(new_user)
}