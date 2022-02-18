package random

import (
	"math/rand"
)

//
func Bool() bool {
	r := rand.Intn(2) // [0,2[
	return r == 0
}

// IntMinMax returns an integer included in the range [min,max]
// it will panic if max-min <= 0
func IntMinMax(min, max int) int {
	r := rand.Intn(max-min+1) + min
	return r
}

//
func IntAround(val int, amplitude int) int {
	r := IntMinMax(val-(amplitude/2), val+(amplitude/2))
	return r
}
