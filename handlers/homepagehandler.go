package handlers

import (
	"fmt"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the betting home page!")
}
