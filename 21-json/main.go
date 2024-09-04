package main

import (
	"encoding/json"
	"fmt"
)

type transaction struct {
	UserId      int      `json:"user_id"` // aliases
	Value       float64  `json:"value"`
	Description string   `json:"description"`
	Hash        string   `json:"-"`              // will be removed from json
	Tags        []string `json:"tags,omitempty"` // will omit nil values
}

func main() {
	fmt.Println("Welcome to JSON study")

	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	transactions := []transaction{
		{1, 23.0, "Supermarket ABC", "df1513g36", []string{"Food", "Groceries"}},
		{2, 154.9, "Steam", "nn6j423n62", nil},
		{3, 15.2, "Drugstore", "47646k24l6", []string{"Medicine", "Healthcare"}},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(transactions, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "user_id": 3,
                "value": 15.2,
                "description": "Drugstore",
                "tags": ["Medicine","Healthcare"]
        }
	`)

	var myTransaction transaction 

	jsonIsValid := json.Valid(jsonDataFromWeb)

	if jsonIsValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &myTransaction)

		fmt.Printf("%#v", myTransaction)
		fmt.Println("Transaction value: ", myTransaction.Value)
	} else {
		fmt.Println("Invalid JSON")
	}

	// some cases where you just want to add to key value

	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myData)

	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key is %v and values is %v and type is %T\n", k, v, v)
	}
}
