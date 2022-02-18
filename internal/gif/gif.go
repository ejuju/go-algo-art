package gif

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"os"

	"golang.org/x/image/draw"
)

const (
	defaultWidth  = 600
	defaultHeight = 600
)

//
func NewFromImages(delay int, imgs ...image.Image) (*gif.GIF, error) {
	fmt.Printf("making new gif from %v images\n", len(imgs))
	result := &gif.GIF{}
	for i, img := range imgs {
		// resize images
		resized := image.NewRGBA(image.Rect(0, 0, defaultWidth, defaultHeight))
		draw.NearestNeighbor.Scale(resized, resized.Rect, img, img.Bounds(), draw.Over, nil)
		// add gif image to result
		var err error
		result, err = AddImage(result, resized, delay)
		if err != nil {
			return nil, err
		}
		fmt.Println("appended img to gif", i, "/", len(imgs))
	}
	fmt.Println("images in gif:", len(result.Image))
	return result, nil
}

//
func AddImage(g *gif.GIF, img image.Image, delay int) (*gif.GIF, error) {
	paletted, err := EncodeImage(img) // convert regular image to paletted image for gif
	if err != nil {
		return nil, err
	}
	g.Delay = append(g.Delay, delay)
	g.Image = append(g.Image, paletted)
	return g, nil
}

//
func EncodeImage(img image.Image) (*image.Paletted, error) {
	// encode image in gif format in bytes buffer
	buf := bytes.Buffer{}
	err := gif.Encode(&buf, img, nil)
	if err != nil {
		return nil, err
	}

	// decode buffer to gif type
	decodedImg, err := gif.Decode(&buf)
	if err != nil {
		return nil, err
	}

	// assert to paletted image
	palettedImg, ok := decodedImg.(*image.Paletted)
	if ok == false {
		return nil, errors.New("unable to assert type image.Image to image.Paletted")
	}

	return palettedImg, nil
}

//
func Save(g *gif.GIF, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if g == nil {
		return errors.New("No GIF was provided (empty pointer)")
	}
	return gif.EncodeAll(file, g)
}
