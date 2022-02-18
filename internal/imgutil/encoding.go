package imgutil

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

// OpenAndDecode opens the file for the given path and returns the decoded image file, the image type, and an error
func OpenAndDecode(path string) (image.Image, string, error) {
	var result image.Image

	// open image file
	imgFile, err := os.Open(path)
	if err != nil {
		log.Println(err.Error())
		return result, "", err
	}
	defer imgFile.Close()

	// decode image
	img, typ, err := image.Decode(imgFile)
	return img, typ, err
}

// SavePNG creates a new file at the given path on your local disk and encodes the provided image inside in the PNG format
func SavePNG(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

// SaveJPG creates a new file at the given path on your local disk and encodes the provided image inside in the JPEG format
func SaveJPG(path string, img image.Image, opts *jpeg.Options) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, img, opts)
}
