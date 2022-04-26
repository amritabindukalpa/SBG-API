package modelcreators

import "github.com/amritabindukalpa/sbg-api/models"

func CreateResult() models.Result {
	result := models.Result{}
	result.SetNumber()
	result.SetColour()
	result.SetParity()
	return result
}
