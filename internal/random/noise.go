package random

import (
	"github.com/josebalius/go-noise"
)

// Perlin2D generates random noise that is consistent with a 2D surface
// Good value for zoom is 100
func Perlin2D(x, y int, min, max float64, zoom float64) int {
	n := noise.NewNoise().Perlin2(float64(x)/zoom, float64(y)/zoom) // between -1 and 1
	mapped := ((1+n)/2)*(max-min) + min
	result := int(mapped)
	return result
}
