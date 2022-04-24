package models

import (
	"sbg-api/determinator"
	"sbg-api/generators"
)

type Result struct {
	Number int
	Colour string
	IsOdd  bool
}

func NewResult() Result {
	n := generators.RandomNumber()
	r := Result{
		Number: n,
		Colour: determinator.Colour(n),
		IsOdd:  determinator.IsOdd(n),
	}
	return r
}
