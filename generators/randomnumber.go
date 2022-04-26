package generators

import (
	"math/rand"
	"time"
)

func RandomNumber(limit int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(limit)
	return randomNumber
}
