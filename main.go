package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // jpeg init
	_ "image/png"  // png init
	"strconv"
	"time"

	"github.com/ejuju/go-algo-art/pkg/imgutil"
)

func main() {
	start := time.Now()

	routines := []string{
		"assets/0.jpg",
		"assets/1.jpg",
		"assets/2.jpg",
		"assets/3.jpg",
		"assets/4.jpg",
		"assets/5.jpg",
	}

	errorsChan := make(chan error, len(routines))

	for i, input := range routines {
		go work(errorsChan, i, input)
	}

	for waitI := 0; waitI < len(routines); waitI++ {
		err := <-errorsChan
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Done in", time.Now().Sub(start).Seconds(), "seconds")
}

func work(ch chan error, i int, inputPath string) {
	var img image.Image
	var err error

	if inputPath != "" {
		img, _, err = imgutil.OpenAndDecode(inputPath)
	}

	outputPath := "output/output_" + strconv.Itoa(int(time.Now().Unix())) + "_" + strconv.Itoa(i) + ".jpg"
	start := time.Now()

	width := imgutil.WidthFromImg(img)
	height := imgutil.HeightFromImg(img)
	paddingX := width / 4
	paddingY := height / 4

	err = imgutil.
		NewImage(time.Now().String(), img).
		BendPixels(1, paddingX, width-paddingX, paddingY, height-paddingY).
		SaveJPG(outputPath, 100)

	fmt.Println("Routine lasted", time.Now().Sub(start).Seconds(), "seconds")

	ch <- err
}
