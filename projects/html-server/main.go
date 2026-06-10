package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

const PORT = 8080

var filesMap = map[string]string{
	"modals-it-computing-b2": "./files/modals_it_computing_B2.html",
}

func main() {
	addr := fmt.Sprintf("0.0.0.0:%d", PORT)
	fmt.Printf("Running the HTML server on %s\n", addr)

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/{id}", fileHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(addr, r))
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	fileBytes := readFile("./static/index.html")
	if fileBytes == nil {
		w.Write([]byte("File not found :("))
	}
	fileContent := string(fileBytes)
	var finalFileListHTML strings.Builder

	for fileName, filePath := range filesMap {
		fmt.Fprintf(&finalFileListHTML, "<li><a href=\"/%s\"><span class=\"file-id\">%s</span><span class=\"file-path\">%s</span></a></li>", fileName, fileName, filePath)
	}

	fileContentWithFilesList := strings.Replace(fileContent, "{% files_list %}", finalFileListHTML.String(), 1)

	w.Write([]byte(fileContentWithFilesList))
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fileBytes := readFile(filesMap[id])
	if fileBytes == nil {
		w.Write([]byte("File not found :("))
	}
	w.Write(fileBytes)
}

func readFile(filename string) []byte {
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return dataByte
}
