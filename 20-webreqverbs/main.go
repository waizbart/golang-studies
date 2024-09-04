package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verbs study")

	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormDataRequest()
}

func PerformGetRequest() {
	const myurl = "https://x-colors.yurace.pro/api/random" // api for random color

	res, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status code: ", res.Status)
	fmt.Println("Content length is: ", res.ContentLength)

	content, _ := io.ReadAll(res.Body)

	// fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount, responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"userId": 1,
			"id": 1,
			"title": "Some title",
			"body": "Teste body"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))
}

func PerformPostFormDataRequest()  {
	const myurl = "https://jsonplaceholder.typicode.com/posts"

	// form data

	data := url.Values{}
	data.Add("userId", "1")
	data.Add("title", "Form title")
	data.Add("body", "Form body")

	res, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))
}
