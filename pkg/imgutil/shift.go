package imgutil

import (
	"fmt"
	"image"
	"math/rand"

	"github.com/aquilax/go-perlin"
)

//
func ShiftPixels(img image.Image, intensity float64, minX, maxX, minY, maxY int) image.Image {
	fmt.Println("Start pixel shift", Height(img)*Width(img), "pixels")

	noisegen := perlin.NewPerlin(2, 2, 4, rand.Int63())
	imgW := Width(img)
	imgH := Height(img)
	result := image.NewRGBA(image.Rect(0, 0, imgW, imgH))

	// middleY := (minY + maxY) / 2

	callbackFunc := func(x int, y int) error {
		// calc distance to closest edge
		// var distToClosestEdge int
		// if y >= middleY {
		// 	distToClosestEdge = maxY - y
		// } else {
		// 	distToClosestEdge = y - minY
		// }

		// calc shift value (offset increases with dist to closest edge)
		xShift := x
		if x <= maxX && x >= minX && y <= maxY && y >= minY {
			noise := noisegen.Noise2D(float64(y+x)/1000, float64(y)/1000)
			xOffset := int(noise * intensity)
			if x+xOffset > imgW {
				xShift = xOffset
			} else if x+xOffset < 0 {
				xShift = imgW + xOffset
			} else {
				xShift += xOffset
			}
			if x == imgW/2 && y == imgH/2 {
				fmt.Println(noise, xOffset, xShift)
			}
		}

		result.Set(x, y, img.At(xShift, y))
		return nil
	}

	_ = ForEachPixel(img, callbackFunc)
	return result
}
