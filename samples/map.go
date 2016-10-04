package main

import (
	"log"
	"net/http"
)

var languages map[string]string

func init() {
	languages = map[string]string{
		"go": "gopher",
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on localhost:8000")

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	languages["Go"] = "Gopher"
}
