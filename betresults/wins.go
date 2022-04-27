package betresults

import (
	"github.com/amritabindukalpa/sbg-api/constants"
	"github.com/amritabindukalpa/sbg-api/models"
)

func GetWins(b models.Bets, r models.Result) models.Winning {
	var w models.Winning
	if b.Number == r.Number {
		w.NumberWin = constants.NumberWin
	}
	if b.Colour == r.Colour {
		w.ColourWin = constants.ColourWin
	}
	if b.IsOdd == r.IsOdd {
		w.ParityWin = constants.ParityWin
	}
	return w
}
