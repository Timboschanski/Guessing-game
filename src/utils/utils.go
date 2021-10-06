package utils

import (
	"math/rand"
	"time"
)

func CutTo(recs [][]string, num int) [][]string {
	lenr := len(recs)
	if lenr > num {
		lenr = num
	}
	recs = recs[:lenr]

	return recs
}

//Randomizer Randomize a number inbetween given parameters
func Randomizer(from int, to int) int {

	rand.Seed(time.Now().UnixNano())
	rdm := (rand.Intn(to-from) + from)
	return rdm
}
