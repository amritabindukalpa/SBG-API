package handlers

import (
	"fmt"
	"net/http"

	"github.com/amritabindukalpa/sbg-api/betresults"
	"github.com/amritabindukalpa/sbg-api/modelcreators"
	"github.com/amritabindukalpa/sbg-api/models"
)

func BetsWithAmount(w http.ResponseWriter, r *http.Request) {
	b1 := models.BetsWithAmount{
		Number: []int{1}, Amount: 2,
	}
	m := map[string]models.BetsWithAmount{
		"Straight up": b1,
	}

	result := modelcreators.CreateResult()
	b := modelcreators.CreateBet(m)
	winning := betresults.GetWins(b, result)
	fmt.Println(winning.ColourWin)
}
