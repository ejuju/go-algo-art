package imgutil

import "image"

// ForEachPixel executes the provided callback for each pixel on a given image
// A position (with X and Y coordinates) as well as the image is provided to the callback function
// The callback function can return an error that will stop the execution of the loop and be returned to the caller
// To access the color from the current pixel, use img.At(x, y)
func ForEachPixel(img image.Image, callback PixelCallbackFunc) error {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			err := callback(x, y)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// PixelCallbackFunc represent a callback function executed for every pixel in the loop of the ForEachPixel function
// It receives the current x and y position within the loop
type PixelCallbackFunc func(x, y int) error
