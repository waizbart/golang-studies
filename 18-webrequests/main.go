package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://waizbart.dev"

func main()  {
	fmt.Println("Web request")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", res)

	defer res.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}