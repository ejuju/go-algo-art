package imgutil

import "image"

//
func (img *Image) Width() int {
	return WidthFromImg(img.Img)
}

//
func (img *Image) Height() int {
	return HeightFromImg(img.Img)
}

func WidthFromImg(img image.Image) int {
	return img.Bounds().Size().X
}

func HeightFromImg(img image.Image) int {
	return img.Bounds().Size().Y
}
