package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://waizbart.dev:3000/projects?tag=backend&lang=golang"

func main()  {
	fmt.Println("Welcome to handling url in golang")
	fmt.Println(myurl)

	// parsing the url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["lang"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "waizbart.dev",
		Path: "/about",
		RawPath: "user=guilherme",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}