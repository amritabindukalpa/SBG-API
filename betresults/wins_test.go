package betresults_test

import (
	"sbg-api/betresults"
	"sbg-api/models"
	"testing"
)

type WinsTestCase struct {
	bet      models.Bets
	result   models.Result
	expected models.Winning
}

var winsTests = []WinsTestCase{
	{
		bet:      models.Bets{2, "red", true},
		result:   models.Result{2, "red", true},
		expected: models.Winning{betresults.NumberWin, betresults.ColourWin, betresults.ParityWin},
	},
	{
		bet:      models.Bets{2, "red", true},
		result:   models.Result{3, "red", true},
		expected: models.Winning{"", betresults.ColourWin, betresults.ParityWin},
	},
	{
		bet:      models.Bets{2, "red", true},
		result:   models.Result{3, "black", false},
		expected: models.Winning{"", "", ""},
	},
	{
		bet:      models.Bets{2, "red", true},
		result:   models.Result{3, "black", false},
		expected: models.Winning{"", "", ""},
	},
}

func TestColourOutput(t *testing.T) {
	for _, tc := range winsTests {
		actual := betresults.GetWins(tc.bet, tc.result)
		if actual.NumberWin != tc.expected.NumberWin {
			t.Fatalf("Number win did not match")
		}
		if actual.ColourWin != tc.expected.ColourWin {
			t.Fatalf("Colour win did not match")
		}
		if actual.ParityWin != tc.expected.ParityWin {
			t.Fatalf("Parity win did not match")
		}
	}
}
