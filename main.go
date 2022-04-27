package main

import (
	"log"
	"net/http"

	"github.com/amritabindukalpa/sbg-api/handlers"
)

func handleRequests() {
	http.HandleFunc("/", handlers.GetHomePage)
	http.HandleFunc("/placebets", handlers.PlaceBets)
	http.HandleFunc("/betswithamount", handlers.BetsWithAmount)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
