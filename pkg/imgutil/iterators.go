package imgutil

import "image"

// ForEachPixel executes the provided callback for each pixel in a given image
// A position (with X and Y coordinates) as well as the ImagePixel is provided to the callback function
// The callback function can return an error that will stop the execution of the loop and be returned to the caller
func ForEachPixel(img image.Image, callback func(x int, y int) error) error {
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
