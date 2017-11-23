package main

import (
	"github.com/djsavvy/lowpoly/blur"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {

	inputImageFilename := os.Args[1]
	reader, err := os.Open(inputImageFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	inputImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	blurredImage, err := blur.GaussianBlur(&inputImage, 1.4, true)
	blurredOutputFile, err := os.Create(inputImageFilename + " blurred.png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(blurredOutputFile, blurredImage)
	if err != nil {
		log.Fatal(err)
	}

}
