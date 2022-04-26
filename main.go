package main

import (
	"log"
	"net/http"

	"github.com/amritabindukalpa/sbg-api/handlers"
)

func handleRequests() {
	http.HandleFunc("/", handlers.PlaceBets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
