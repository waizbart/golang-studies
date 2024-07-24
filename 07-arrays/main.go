package main

import "fmt"

func main()  {
	fmt.Println("Welcome to the array study")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[2] = "Tomato"
	fruitList[3] = "Watermelon"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Number of fruits: ", len(fruitList))

	var carList = [3]string{"Ferrari", "Volkswagen", "Fiat"}

	fmt.Println("Cars: ", carList)
	fmt.Println("Cars: ", len(carList))
	fmt.Println("Second car: ", carList[1])
}