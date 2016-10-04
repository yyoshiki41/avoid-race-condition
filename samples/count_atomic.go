package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var count int64

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Println("Starting server on localhost:8000")

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&count, 1)
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Count %d\n", count)
}
