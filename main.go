package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world\n")
}
