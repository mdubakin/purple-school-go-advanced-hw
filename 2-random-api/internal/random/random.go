package random

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func getRandomInt(max int) int {
	return r.Intn(max + 1)
}
