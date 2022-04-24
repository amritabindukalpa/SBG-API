package handlers

import (
	"fmt"
	"net/http"
	"sbg-api/betresults"
	"sbg-api/models"
	"strconv"
)

func PlaceBets(w http.ResponseWriter, r *http.Request) {
	var odd bool
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	c := r.URL.Query().Get("c")
	p := r.URL.Query().Get("p")
	if p == "1" {
		odd = true
	}
	b := models.Bets{
		Number: n,
		Colour: c,
		IsOdd:  odd,
	}

	result := models.NewResult()
	wins := betresults.GetWins(b, result)
	fmt.Fprintf(w, wins.NumberWin+wins.ColourWin+wins.ParityWin+"The number is:"+strconv.Itoa(result.Number))
}
