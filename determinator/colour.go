package determinator

func Colour(randomNumber int) string {

	if (1 <= randomNumber && randomNumber <= 10) || (19 <= randomNumber && randomNumber <= 28) {
		if IsOdd(randomNumber) {
			return "red"
		}
		return "black"
	}
	if IsOdd(randomNumber) {
		return "black"
	}
	return "red"
}