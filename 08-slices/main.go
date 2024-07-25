package main

import (
	"fmt"
	// "sort"
)

func main() {
	fmt.Println("Welcome to my study about go slices")

	var top5BestNumbers = []int{7, 9, 13}
	// fmt.Printf("Type of top list is: %T\n", top5BestNumbers)

	top5BestNumbers = append(top5BestNumbers, 2004, 17)
	// fmt.Println("List size: ", len(top5BestNumbers))

	top5BestNumbers = top5BestNumbers[2:4]
	// fmt.Println("List", top5BestNumbers)

	highScore := make([]int, 4)

	highScore[0] = 30
	highScore[1] = 52
	highScore[2] = 24
	highScore[3] = 16
	// highScore[4] = 77

	highScore = append(highScore, 10, 12, 43)

	fmt.Println(highScore)

	// sort.Ints(highScore)

	// fmt.Println("SORTED: ", highScore)
	// fmt.Println("IS SORTED: ", sort.IntsAreSorted(highScore))

	// how to remove a value from slice based on index

	var courses = []string{"Python", "React", "GO", "Javascript", "Ruby"}

	fmt.Println(courses)

	index := 2

	courses = append(courses[:index], courses[index+1:]...)
	
	fmt.Println(courses)
}