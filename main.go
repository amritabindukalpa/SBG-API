package main

import (
	"log"
	"net/http"
	"sbg-api/handlers"
)

func handleRequests() {
	http.HandleFunc("/", handlers.PlaceBets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
