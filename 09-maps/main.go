package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	
	languages["py"] = "Python"
	languages["js"] = "Javascript"
	languages["html"] = "Hyper text markup language"

	fmt.Println(languages)
	fmt.Println(languages["html"])
	fmt.Println(languages["ts"])

	delete(languages, "js")

	fmt.Println(languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("KEY: %v  VALUE: %v\n", key, value)
	}

	var numbers = [3]int{1, 2, 3}

	for i, v := range numbers {
		fmt.Printf("%d: %d\n", i, v)
	}
}