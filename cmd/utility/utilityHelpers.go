package utility

import (
	"math/rand"
	"strings"
	"time"
	"github.com/rubensseva/timp/cmd/data/model"
)

// RandomGen generates a random number between range
func RandomGen(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func CalcWPM(text model.Text, elapsed time.Duration) float32 {
	totNumOfWords := len(strings.Fields(text.GetText()))
	var wpm float32 = float32(totNumOfWords) / float32(float32(elapsed.Seconds())/60.0)
	return wpm
}
