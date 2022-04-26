package models

import (
	"github.com/amritabindukalpa/sbg-api/constants"
	"github.com/amritabindukalpa/sbg-api/determinator"
	"github.com/amritabindukalpa/sbg-api/generators"
)

type Result struct {
	Number int
	Colour string
	IsOdd  bool
}

func NewResult() Result {
	n := generators.RandomNumber(constants.NumberLimit)
	r := Result{
		Number: n,
		Colour: determinator.Colour(n),
		IsOdd:  determinator.IsOdd(n),
	}
	return r
}
