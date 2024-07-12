package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	fmt.Println("Present raw time: ", presentTime)
	fmt.Println("Present formatted time: ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2004, time.September, 13, 0, 0, 0, 0, time.UTC)
	fmt.Println("Created raw date: ", createdDate)
	fmt.Println("Created formatted date: ", createdDate.Format("Year: 2006, Month: 01, Day: 02, Week day: Monday"))
}