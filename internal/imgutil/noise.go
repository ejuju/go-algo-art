package imgutil

import (
	"image"
	"math"

	"github.com/ejuju/go-algo-art/internal/random"
)

// ApplyNoise ...
// Good value for zoom is 100
func ApplyNoise(maxIntensity float64, zoom float64) TransformationFunc {
	return func(img image.Image) image.Image {
		// minX, minY, maxX, maxY := 0, 0, Width(img), Height(img)
		minX := Width(img) / 4
		minY := Height(img) / 4
		maxX := 3 * Width(img) / 4
		maxY := 3 * Height(img) / 4

		output := image.NewRGBA(image.Rect(0, 0, Width(img), Height(img)))

		ForEachPixel(img, func(x, y int) error {
			xShift := 0
			yShift := 0

			if x > minX && x < maxX && y > minY && y < maxY {
				intensity := math.Pow(float64(getDistToClosestEdge(x, y, minX, minY, maxX, maxY))*0.1, 1.1)

				if intensity > maxIntensity {
					intensity = maxIntensity
				}

				xShift = random.Perlin2D(x, y, -intensity, intensity, zoom)
				yShift = random.Perlin2D(x+31, y-13, -intensity, intensity, zoom)

				if x+xShift > Width(img) || x+xShift < 0 {
					xShift = 0
				}

				if y+yShift > Height(img) || y+yShift < 0 {
					yShift = 0
				}
			}

			output.Set(x, y, img.At(x+xShift, y+yShift))
			return nil
		})

		return output
	}
}

func getDistToClosestEdge(x, y int, minX, minY, maxX, maxY int) int {
	result := 0

	distToMinX := int(math.Abs(float64(minX - x)))
	distToMinY := int(math.Abs(float64(minY - y)))

	distToMaxX := int(math.Abs(float64(maxX - x)))
	distToMaxY := int(math.Abs(float64(maxY - y)))

	if distToMinX > result {
		result = distToMinX
	}

	if distToMinY > distToMinX {
		result = distToMinY
	}

	if distToMaxX > distToMinY {
		result = distToMaxX
	}

	if distToMaxY > distToMaxX {
		result = distToMaxY
	}

	return result
}
