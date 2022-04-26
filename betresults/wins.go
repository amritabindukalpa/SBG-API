package betresults

import (
	"github.com/amritabindukalpa/sbg-api/models"
)

const NumberWin = "Your number won!"
const ColourWin = "Your colour won!"
const ParityWin = "Your parity won!"

func GetWins(b models.Bets, r models.Result) models.Winning {
	var w models.Winning
	if b.Number == r.Number {
		w.NumberWin = NumberWin
	}
	if b.Colour == r.Colour {
		w.ColourWin = ColourWin
	}
	if b.IsOdd == r.IsOdd {
		w.ParityWin = ParityWin
	}
	return w
}
