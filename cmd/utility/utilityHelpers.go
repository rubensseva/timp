package utility

import (
	"math/rand"
	"time"
)

// RandomGen generates a random number between range
func RandomGen(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
