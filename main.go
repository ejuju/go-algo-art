package main

import (
	"fmt"
	_ "image/jpeg" // jpeg init
	_ "image/png"  // png init
	"log"
	"time"

	"github.com/ejuju/go-algo-art/pkg/imgutil"
)

func main() {
	const inputPath = "assets/input_img.jpg"
	const outputPath = "output/output.png"
	img, _, err := imgutil.OpenAndDecode(inputPath)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Pixel shifting
	start := time.Now()

	// xPadding := 0
	xPadding := imgutil.Width(img) / 3
	yPadding := imgutil.Height(img) / 4
	pixelShifted := imgutil.ShiftPixels(
		img,
		2500,
		xPadding,
		imgutil.Width(img)-xPadding,
		yPadding,
		imgutil.Height(img)-yPadding,
	)

	fmt.Println("Finished pixel shifting", "Done in", time.Now().Sub(start).Seconds(), "seconds")

	// save image as png
	err = imgutil.EncodeToNewFileWithPNG(outputPath, pixelShifted)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println("Saved output", "Done after", time.Now().Sub(start).Seconds(), "seconds")
}
