package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amritabindukalpa/sbg-api/betresults"
	"github.com/amritabindukalpa/sbg-api/modelcreators"
	"github.com/amritabindukalpa/sbg-api/models"
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

	result := modelcreators.CreateResult()
	wins := betresults.GetWins(b, result)
	fmt.Fprintf(w, wins.NumberWin+wins.ColourWin+wins.ParityWin+"The winning number is:"+strconv.Itoa(result.Number))
}
