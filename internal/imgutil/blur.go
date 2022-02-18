package imgutil

import (
	"image"
	"math/rand"
)

//
func ApplyBlur() TransformationFunc {
	return func(img image.Image) image.Image {
		newImg := image.NewRGBA(image.Rect(0, 0, Width(img), Height(img)))

		minX := Width(img) / 4
		maxX := 3 * Width(img) / 4
		minY := Height(img) / 4
		maxY := 3 * Height(img) / 4

		_ = ForEachPixel(img, func(x, y int) error {
			xShift := 0
			yShift := 0

			randAmpl := 50

			if x > minX && x < maxX && y > minY && y < maxY {
				xShift = rand.Intn(randAmpl) - (randAmpl / 2)
				yShift = rand.Intn(randAmpl) - (randAmpl / 2)
			}

			newImg.Set(x, y, img.At(x+xShift, y+yShift))
			return nil
		})

		return newImg
	}
}

//
func BlurPixels(img image.Image) image.Image {
	newImg := image.NewRGBA(image.Rect(0, 0, Width(img), Height(img)))

	minX := Width(img) / 4
	maxX := 3 * Width(img) / 4
	minY := Height(img) / 4
	maxY := 3 * Height(img) / 4

	_ = ForEachPixel(img, func(x, y int) error {
		xShift := 0
		yShift := 0

		randAmpl := 50

		if x > minX && x < maxX && y > minY && y < maxY {
			xShift = rand.Intn(randAmpl) - (randAmpl / 2)
			yShift = rand.Intn(randAmpl) - (randAmpl / 2)
		}

		newImg.Set(x, y, img.At(x+xShift, y+yShift))
		return nil
	})

	return newImg
}
