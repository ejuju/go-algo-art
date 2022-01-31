package imgutil

import (
	"image"
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

// EncodeToNewFileWithPNG creates a new file at the given path on your local disk and encodes the provided image inside
// This function uses the png.Encode() function from the standard library
func EncodeToNewFileWithPNG(path string, img image.Image) error {
	// Create a empty file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// encode and return error
	return png.Encode(file, img)
}
