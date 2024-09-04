package main

import (
	"fmt"
	"net/http"
    "log"

	"github.com/gorilla/mux"
)

func main() {
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":3333", r))
}

func greeter() {
	fmt.Println("Hey there mod users!!!")
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("<h1>Hello everybody</h1>"))
}
