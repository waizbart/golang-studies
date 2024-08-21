package main

import "fmt"

func main() {
	fmt.Println("Welcome to the go loops study")

	// days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for i, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", i, day)
	// }

	rogueValue := 1

	// for rogueValue < 10 {
	// 	fmt.Println("Value is: ", rogueValue)
	// 	rogueValue++
	// }

	for {
		if rogueValue == 10 {
			break
		}

		if rogueValue == 2 {
			goto lco
		}

		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping hereee")
	
}
