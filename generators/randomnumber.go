package generators

import (
	"math/rand"
	"time"
)

func RandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(36)
	return randomNumber
}
