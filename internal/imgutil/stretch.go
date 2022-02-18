package imgutil

import (
	"image"

	"golang.org/x/image/draw"
)

//
func ApplyStretch(width, height int) TransformationFunc {
	return func(img image.Image) image.Image {
		output := image.NewRGBA(image.Rect(0, 0, width, height))
		draw.NearestNeighbor.Scale(output, output.Rect, img, img.Bounds(), draw.Over, nil)
		return output
	}
}
