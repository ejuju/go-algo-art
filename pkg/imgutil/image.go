package imgutil

import "image"

//
type Image struct {
	Title string
	Img   image.Image
}

//
func NewImage(title string, img image.Image) *Image {
	return &Image{
		Title: title,
		Img:   img,
	}
}
