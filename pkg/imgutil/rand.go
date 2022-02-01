package imgutil

import (
	"image"
	"image/color"
	"math"

	"github.com/aquilax/go-perlin"
)

//
func (img *Image) Rand(width int, height int, zoom float64, seed int64) *Image {
	noisegen := perlin.NewPerlin(2, 2, 3, seed)

	result := image.NewRGBA(image.Rect(0, 0, width, height))

	if zoom < 20 {
		zoom = 20
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			val := noiseToRGBA(noisegen.Noise2D(intToNoiseInput(x, zoom), intToNoiseInput(y, zoom)))
			c := color.RGBA{
				R: val,
				G: val,
				B: val,
				A: 255,
			}

			result.Set(x, y, c)
		}
	}

	img.Img = result
	return img
}

func intToNoiseInput(i int, zoom float64) float64 {
	return float64(i) / zoom
}

func noiseToRGBA(noise float64) uint8 {
	normN := math.Abs((.5 + noise) * 255)
	uin := uint8(normN)
	if uin < 210 {
		return 0
	}
	return uin
}
