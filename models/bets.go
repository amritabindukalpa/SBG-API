package models

//import "github.com/go-playground/validator"

type Bets struct {
	Number int    //`validate :"required,gte=20,lte=65"`
	Colour string //`validate: ="red", ="black"`
	IsOdd  bool
}

// var validate *validator.Validate
