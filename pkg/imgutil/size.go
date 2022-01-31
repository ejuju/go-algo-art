package imgutil

import "image"

//
func Width(img image.Image) int {
	return img.Bounds().Size().X
}

//
func Height(img image.Image) int {
	return img.Bounds().Size().Y
}
