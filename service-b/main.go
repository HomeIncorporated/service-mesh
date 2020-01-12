package main

import (
	"fmt"
	"log"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from service B  ")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("I'm running...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8788", nil))
}