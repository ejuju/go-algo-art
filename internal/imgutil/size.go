package imgutil

import "image"

//
func Width(img image.Image) int {
	return img.Bounds().Dx()
}

//
func Height(img image.Image) int {
	return img.Bounds().Dy()
}
