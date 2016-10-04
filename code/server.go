package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes message
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, gophers")
}
