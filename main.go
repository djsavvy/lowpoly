package main

import (
	"fmt"
	"github.com/djsavvy/lowpoly/average"
	"github.com/djsavvy/lowpoly/blur"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {

	for _, inputImageFilename := range os.Args[1:] {

		reader, err := os.Open(inputImageFilename)
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()

		inputImage, _, err := image.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}

		//		blurTester(&inputImage, inputImageFilename)

		averageTester(&inputImage, inputImageFilename)
	}
}

func blurTester(inputImage *image.Image, inputImageFilename string) {

	fmt.Println("testing blur")

	blurredImage, err := blur.GaussianBlur(inputImage, 15, true)
	if err != nil {
		log.Fatal(err)
	}
	blurredOutputFile, err := os.Create(inputImageFilename + " blurred.png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(blurredOutputFile, blurredImage)
	if err != nil {
		log.Fatal(err)
	}
}

func averageTester(inputImage *image.Image, inputImageFilename string) {

	fmt.Println("testing average")

	output := image.NewRGBA((*inputImage).Bounds())
	err := average.TriangleAverage(inputImage, output, &image.Point{100, 100}, &image.Point{500, 300}, &image.Point{200, 800})
	if err != nil {
		log.Fatal(err)
	}
	averagedOutputFile, err := os.Create(inputImageFilename + " averaged.png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(averagedOutputFile, output)
	if err != nil {
		log.Fatal(err)
	}
}
