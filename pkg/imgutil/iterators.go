package imgutil

import (
	"image/color"
)

// PixelCallbackFunc represents a function that can be executed on a pixel
// It receives the x and y positions of the pixel on the image and the color of the current pixel
type PixelCallbackFunc func(x int, y int, c color.Color, args ...interface{}) error

// ForEachPixel executes the provided callback for each pixel in a given image
// A position (with X and Y coordinates) as well as the ImagePixel is provided to the callback function
// The callback function can return an error that will stop the execution of the loop and be returned to the caller
func (img *Image) ForEachPixel(callback PixelCallbackFunc, args ...interface{}) error {
	bounds := img.Img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			err := callback(x, y, img.Img.At(x, y), args...)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
