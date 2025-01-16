package random

import (
	"math/rand"
	"time"
)

func getRandomInt(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	return r.Intn(max + 1)
}
