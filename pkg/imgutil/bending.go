package imgutil

import (
	"image"
	"math/rand"

	"github.com/aquilax/go-perlin"
)

//
func (img *Image) BendPixels(intensity float64, minX, maxX, minY, maxY int) *Image {
	noisegen := perlin.NewPerlin(2, 2, 4, rand.Int63())
	imgW := img.Width()
	imgH := img.Height()
	result := image.NewRGBA(image.Rect(0, 0, imgW, imgH))

	for y := 0; y < imgH; y++ {
		for x := 0; x < imgW; x++ {
			if x > minX && x < maxX && y > minY && y < maxY {
				xShift := x
				yShift := x

				// noise := noisegen.Noise2D(float64(x)/1000, float64(y)/1000)
				xNoise := noisegen.Noise2D(float64(x+2)/1000, float64(y)/1000)
				yNoise := noisegen.Noise2D(float64(x+1)/1000, float64(y+4)/1000)

				distToClosestEdge := calcDistToClosestEdge(x, y, minX, maxX, minY, maxY)

				xOffset := calcNoiseToOffset(xNoise, distToClosestEdge)
				yOffset := calcNoiseToOffset(yNoise, distToClosestEdge)

				xShift += xOffset
				yShift += yOffset

				result.Set(x, y, img.Img.At(xShift, yShift))
			} else {
				result.Set(x, y, img.Img.At(x, y))
			}
		}
	}

	img.Img = result
	return img
}

func calcNoiseToOffset(noise float64, distToClosestEdge int) int {
	result := noise * float64(distToClosestEdge) / 1000
	return int(result)
}

func calcDistToClosestEdge(x, y, minX, maxX, minY, maxY int) int {
	var lowestDist int
	loop := []int{
		x - minX,
		maxX - x,
		y - minY,
		maxY - y,
	}
	for _, val := range loop {
		if val > lowestDist {
			lowestDist = val
		}
	}
	return lowestDist
}
