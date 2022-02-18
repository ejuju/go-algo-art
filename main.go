package main

import (
	"fmt"
	_ "image/jpeg" // jpeg init
	_ "image/png"  // png init
	"strconv"
	"time"

	"github.com/ejuju/go-algo-art/internal/imgutil"
)

func main() {
	// rand.Seed(time.Now().UnixNano())

	start := time.Now()
	fmt.Println("Launching...")

	// assets := []string{
	// 	"assets/0.jpg",
	// 	"assets/1.jpg",
	// 	"assets/2.jpg",
	// 	"assets/4.jpg",
	// 	"assets/5.jpg",
	// }

	// run := func(id string, assets []string) error {
	// 	// define output file path
	// 	outputPrefix := "output/output_"
	// 	outputFmt := ".png"
	// 	outputPath := outputPrefix + strconv.Itoa(int(time.Now().Unix())) + "_" + id + outputFmt

	// 	// open input img
	// 	img, _, err := imgutil.OpenAndDecode(assets[0])
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// transform img
	// 	outputImg := imgutil.Transform(
	// 		img,
	// 		imgutil.ApplyNoise(1000, 200),
	// 	)

	// 	err = imgutil.SavePNG(outputPath, outputImg) // save result

	// 	return err
	// }

	// w1 := worker.NewWorker("imgnoise",
	// 	worker.NewTask("bend", assets, run),
	// )
	// errs, dur := w1.Work()
	// fmt.Println(errs, dur)

	// jobs := assets

	// for i := range jobs {
	// 	go work(errorsChan, i, assets)
	// }

	// for waitI := 0; waitI < len(jobs); waitI++ {
	// 	err := <-errorsChan
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	fmt.Println("Done in", time.Now().Sub(start).Seconds(), "seconds")
}

func work(ch chan error, jobID int, assets []string) {
	// start timer
	start := time.Now()
	fmt.Println("Starting job", jobID)

	// define output file path
	outputPrefix := "output/output_"
	outputFmt := ".png"
	outputPath := outputPrefix + strconv.Itoa(int(time.Now().Unix())) + "_" + strconv.Itoa(jobID) + outputFmt

	// open input img
	img, _, err := imgutil.OpenAndDecode(assets[jobID])
	if err != nil {
		ch <- err
		return
	}

	// transform img
	outputImg := imgutil.Transform(
		img,
		imgutil.ApplyNoise(1000, 200),
	)

	err = imgutil.SavePNG(outputPath, outputImg) // save result

	fmt.Println("error:", err)

	fmt.Println("Job took", time.Now().Sub(start).Seconds(), "seconds", "[jobID", jobID, "]")

	ch <- err
}
