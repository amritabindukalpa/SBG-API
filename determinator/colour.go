package determinator

import (
	"github.com/amritabindukalpa/sbg-api/constants"
)

func Colour(randomNumber int) string {

	if (1 <= randomNumber && randomNumber <= 10) || (19 <= randomNumber && randomNumber <= 28) {
		if IsOdd(randomNumber) {
			return constants.Red
		}
		return constants.Black
	}
	if IsOdd(randomNumber) {
		return constants.Black
	}
	return constants.Red
}
