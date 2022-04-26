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

func (f *Result) SetNumber() {
    f.Number = generators.RandomNumber(constants.NumberLimit)
}

func (f *Result) SetColour() {
    f.Colour = determinator.Colour(f.Number)
}

func (f *Result) SetParity() {
    f.IsOdd = determinator.IsOdd(f.Number)
}
